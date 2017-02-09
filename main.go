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

	"github.com/tschf/odl/arch"
	"github.com/tschf/odl/resource/finder"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/crypto/ssh/terminal"
)

func checkRedirect(req *http.Request, via []*http.Request) error {
	req.Header.Add("User-Agent", "Mozilla/5.0")

	return nil
}

// In order to download software, the OTN license agreement must be accepted.
// This is implemented by two mechanisms:
// 1. The command line argument `accept-license`
// 2. Prompting for user input, with a Y/N value where Y indicated truth
func getLicenseAcceptance(acceptFromFlag bool, licenseURL string) bool {

	if acceptFromFlag {
		return true
	}

	// Because the accept flag wasn't passed in, we want to prompt the user
	// to decide if they'd like to accept the license or not - passing in the
	// URL to the license agreement
	fmt.Println("Do you accept the OTN license agreement?")
	fmt.Println(fmt.Sprintf("Full terms found here: %s", licenseURL))
	fmt.Print("Enter Y for Yes, or N for No: ")

	reader := bufio.NewReader(os.Stdin)
	strLicenseAccepted, _ := reader.ReadString('\n')

	return strings.TrimSpace(strLicenseAccepted) == "Y"
}

func main() {

	var (
		flagUser          = flag.String("username", "", "Specify the user account that will be logging in and accepting the license agreement. Alternatively, set the environment variable OTN_USERNAME.")
		flagPassword      = flag.String("password", "", "Specify the password that corresponds to your OTN account. Alternatively, set the environment variable OTN_PASSWORD.")
		flagOs            = flag.String("os", "linux", "Specify the desired platform of the software. Should be \"linux\" or \"windows\"")
		flagComponent     = flag.String("component", "", "Specify the component to grab.")
		flagVersion       = flag.String("version", "", "Specify the software version. ")
		flagLang          = flag.String("lang", "na", "Specify the language of the software. Should be \"en\" or \"na\"")
		flagAcceptLicense = flag.Bool("accept-license", false, "Specify whether or not you accept the OTN license agreement for the nominated software.")
		flagArchitecture  arch.Arch
	)

	flag.Var(&flagArchitecture, "arch", "Specify the desired architecture of the software. Should be \"x86\", \"x64\", or \"na\"")
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

	selectedFile, ok := finder.FindResource(*flagComponent, *flagVersion, *flagOs, flagArchitecture, *flagLang)

	if ok {

		req, _ := http.NewRequest("GET", selectedFile.File, nil)
		req.Header.Add("User-Agent", "Mozilla/5.0")

		// The license accepted is done either through a command line flag
		// (accept-license), or if that is not set, prompted the user for input
		// which accepts a Y/N value. Only if the user inputs Y does that
		// indicate acceptance of the license.

		otnLicenseAccepted := getLicenseAcceptance(*flagAcceptLicense, selectedFile.License)

		if !otnLicenseAccepted {
			fmt.Fprint(os.Stderr, "You must accept the license agreement in order to download. Exiting now.\n")
			os.Exit(1)
		}

		var cookies []*http.Cookie
		cookies = append(cookies, selectedFile.AcceptCookie)

		u, _ := url.Parse(selectedFile.File)
		cookieJar, _ := cookiejar.New(nil)

		//Set initial jar with a cookie accepting the license agreement
		//Example taken from: https://gist.github.com/Rabbit52/a8a44c3c4cd514052952
		cookieJar.SetCookies(u, cookies)

		client := &http.Client{
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

		if !selectedFile.SkipAuth {

			if len(otnPassword) == 0 {
				fmt.Printf("Enter your OTN password (%s):", otnUser)
				consolePass, _ := terminal.ReadPassword(int(syscall.Stdin))

				otnPassword = string(consolePass)
			}

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
		}

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
