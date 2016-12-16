package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path"
	"strings"
	"syscall"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/crypto/ssh/terminal"
)

func checkRedirect(req *http.Request, via []*http.Request) error {
	req.Header.Add("User-Agent", "Mozilla/5.0")

	return nil
}

func main() {

	flagDlOption := flag.String("getFile", "", "Specify the file you want to download. Supported files include: db11gxe-linux")
	flagUser := flag.String("username", "", "Specify the user account that will be logging in and accepting the license agreement. Alternatively, set the environment variable OTN_USERNAME.")
	flagPassword := flag.String("password", "", "Specify the password that corresponds to your OTN account. Alternatively, set the environment variable OTN_PASSWORD.")
	flag.Parse()

	otnUser := *flagUser
	if len(otnUser) == 0 {
		otnUser = os.Getenv("OTN_USERNAME")
	}
	//
	fmt.Println(otnUser)
	//
	if len(otnUser) == 0 {
		log.Fatal("You must specify an OTN username to access OTN files. Set with the flag -username or set the environment variable OTN_USERNAME.")
	}
	//
	otnPassword := *flagPassword
	if len(otnPassword) == 0 {
		otnPassword = os.Getenv("OTN_PASSWORD")
	}

	// Following example from: https://gist.github.com/Rabbit52/a8a44c3c4cd514052952
	// Also see: http://stackoverflow.com/questions/30652577/go-doing-a-get-request-and-building-the-querystring

	// Initial prototype. Download Oracle 11g XE (Linux) from the command line
	// Split this out a bit more in the future to add support for more files
	if *flagDlOption == "db11gxe-linux" {

		requestFile := "https://edelivery.oracle.com/akam/otn/linux/oracle11g/xe/oracle-xe-11.2.0-1.0.x86_64.rpm.zip"

		req, _ := http.NewRequest("GET", requestFile, nil)
		req.Header.Add("User-Agent", "Mozilla/5.0")

		// proxyURL, err := url.Parse("http://localhost:8888")
		// if err != nil {
		// 	fmt.Println("Proxy parse fail")
		// 	log.Fatal(err)
		// }
		//
		// tr := &http.Transport{
		// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		// 	Proxy:           http.ProxyURL(proxyURL),
		// }
		// http.DefaultTransport = &http.Transport{
		// 	Proxy: http.ProxyURL(proxyURL),
		// }

		fmt.Println("Do you accept the XE license agreement?")
		fmt.Println("Full terms found here: http://www.oracle.com/technetwork/licenses/database-11g-express-license-459621.html")
		fmt.Print("Enter Y for Yes, or N for No: ")

		reader := bufio.NewReader(os.Stdin)
		licenseAccept, _ := reader.ReadString('\n')

		if strings.TrimSpace(licenseAccept) != "Y" {
			log.Fatal("You must accept the license agreement in order to download. Exiting now.")
		}

		if len(otnPassword) == 0 {
			fmt.Printf("Enter your OTN password (%s):", otnUser)
			consolePass, _ := terminal.ReadPassword(int(syscall.Stdin))

			otnPassword = string(consolePass)
		}

		var cookies []*http.Cookie
		//
		otnAcceptCookie := &http.Cookie{
			Name:   "oraclelicense",
			Value:  "accept-sqldev-cookie",
			Domain: ".oracle.com",
		}
		cookies = append(cookies, otnAcceptCookie)
		u, _ := url.Parse(requestFile)
		cookieJar, _ := cookiejar.New(nil)

		//Set initial jar with a cookie accepting the license agreement
		//Example taken from: https://gist.github.com/Rabbit52/a8a44c3c4cd514052952
		cookieJar.SetCookies(u, cookies)

		client := &http.Client{
			// Transport: tr,
			// Need to re-add user agent as Go doesn't propagate them
			// see: https://groups.google.com/forum/#!topic/golang-nuts/OwGvopYXpwE
			// and https://github.com/golang/go/issues/4800
			// slated to be fixed in golang 1.8. Done in checkRedirect
			CheckRedirect: checkRedirect,
			Jar:           cookieJar,
		}
		resp, respErr := client.Do(req)
		if respErr != nil {
			fmt.Println("Couldnt read resp")
			log.Fatal(respErr)
		}
		defer resp.Body.Close()

		fmt.Printf("The file being requested is %s\n", requestFile)

		doc, err := goquery.NewDocumentFromResponse(resp)
		if err != nil {
			fmt.Println("Error")
			log.Fatal(err)
		}

		// fmt.Println(doc)

		// https://godoc.org/github.com/PuerkitoBio/goquery
		pageForms := doc.Find("form")

		// formAction, _ := pageForms.Attr("action")
		// fmt.Println("POST to>", formAction)

		//POST example: http://stackoverflow.com/questions/19253469/make-a-url-encoded-post-request-using-http-newrequest
		authData := url.Values{}
		pageForms.Find("input").Each(func(index int, el *goquery.Selection) {
			inputName, _ := el.Attr("name")
			inputValue, _ := el.Attr("value")

			authData.Set(inputName, inputValue)
		})

		authData.Set("username", otnUser)
		authData.Set("password", otnPassword)

		req, _ = http.NewRequest("POST", "https://login.oracle.com/oam/server/sso/auth_cred_submit", bytes.NewBufferString(authData.Encode()))
		req.Header.Add("User-Agent", "Mozilla/5.0")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		resp, _ = client.Do(req)
		defer resp.Body.Close()

		file, err := os.Create(path.Base(requestFile))
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Download complete.")
	}

}
