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

	"github.com/tschf/odl/db"
	"github.com/tschf/odl/types"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/crypto/ssh/terminal"
)

func checkRedirect(req *http.Request, via []*http.Request) error {
	req.Header.Add("User-Agent", "Mozilla/5.0")

	return nil
}

func main() {

	flagUser := flag.String("username", "", "Specify the user account that will be logging in and accepting the license agreement. Alternatively, set the environment variable OTN_USERNAME.")
	flagPassword := flag.String("password", "", "Specify the password that corresponds to your OTN account. Alternatively, set the environment variable OTN_PASSWORD.")
	flagOs := flag.String("os", "linux", "Specify the desired platform of the software. Should be \"linux\" or \"windows\"")
	flagComponent := flag.String("component", "db", "Specify the component to grab. Should be \"db\"")
	flagArchitecture := flag.String("architecture", "amd64", "Specify the desired architecture of the software. Should be \"amd64\" or \"32\"")
	flagVersion := flag.String("version", "11gXE", "Specify the software version. Should be \"11gXE\"")

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

	otnPassword := *flagPassword
	if len(otnPassword) == 0 {
		otnPassword = os.Getenv("OTN_PASSWORD")
	}

	//New data structure to store files in, to provide an index system
	//key format will be: "component:os:arch:version"
	var files map[string]*types.Resource
	files = make(map[string]*types.Resource)

	//Get all the files (currently only one type (Oracle 11gXE).)
	xe11gResources := db.GetXeResouces()

	for _, xe11gResource := range xe11gResources {
		files[fmt.Sprintf("%s:%s:%s:%s", xe11gResource.Component, xe11gResource.OS, xe11gResource.Arch, xe11gResource.Version)] = xe11gResource
	}

	// Initial prototype. Download Oracle 11g XE (Linux) from the command line
	// Split this out a bit more in the future to add support for more files
	selectedFile, ok := files[fmt.Sprintf("%s:%s:%s:%s", *flagComponent, *flagOs, *flagArchitecture, *flagVersion)]

	if ok {
		req, _ := http.NewRequest("GET", selectedFile.File, nil)
		req.Header.Add("User-Agent", "Mozilla/5.0")

		fmt.Println("Do you accept the XE license agreement?")
		fmt.Println(fmt.Sprintf("Full terms found here: %s", selectedFile.License))
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
		cookies = append(cookies, selectedFile.AcceptCookie)

		u, _ := url.Parse(selectedFile.File)
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
			fmt.Println("Couldn't read response")
			log.Fatal(respErr)
		}
		defer resp.Body.Close()

		fmt.Printf("The file being requested is %s\n", selectedFile.File)

		doc, err := goquery.NewDocumentFromResponse(resp)
		if err != nil {
			fmt.Println("Error")
			log.Fatal(err)
		}

		// https://godoc.org/github.com/PuerkitoBio/goquery
		pageForms := doc.Find("form")

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

		file, err := os.Create(path.Base(selectedFile.File))
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Download complete.")
	} else {
		log.Fatal("Err, Could not find the selected file.")
	}
}
