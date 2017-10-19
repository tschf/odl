package dl

import (
	"bufio"
	"bytes"
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

	pb "gopkg.in/cheggaaa/pb.v1"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/tschf/odl/resource"
)

func checkRedirect(req *http.Request, via []*http.Request) error {
	req.Header.Add("User-Agent", "Mozilla/5.0")

	return nil
}

func getRequest(method string, path string, requestBody url.Values) *http.Request {
	req, err := http.NewRequest(method, path, bytes.NewBufferString(requestBody.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0")
	if method == http.MethodPost {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	return req
}

func getResponse(req *http.Request, httpClient *http.Client) *http.Response {
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func SaveResource(res *resource.OracleResource, otnUser string, otnPassword string, skipExisting bool) {
	var cookies []*http.Cookie
	var fileResp *http.Response

	cookies = append(cookies, res.AcceptCookie)

	cookieJar, _ := cookiejar.New(nil)

	client := &http.Client{
		// Need to re-add user agent as Go doesn't propagate them
		// see: https://groups.google.com/forum/#!topic/golang-nuts/OwGvopYXpwE
		// and https://github.com/golang/go/issues/4800
		// slated to be fixed in golang 1.8. Done in checkRedirect
		CheckRedirect: checkRedirect,
		Jar:           cookieJar,
	}

	// flag to identify if we need to submit username/password data
	// to https://login.oracle.com/oam/server/sso/auth_cred_submit. Specifically
	// some resources include more than one file (such as Oracle 12c EE).
	// In which case, we only need to submit once.
	// (todo/test: read response headers to see if auth is required?)
	requiresAuth := true

	for _, file := range res.File {
		//Check to see if the file we are requesting to download already exists.
		//Impl borrowed from: http://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
		//err will be nil when getting information about a file that exists
		_, fileStatErr := os.Stat(path.Base(file))
		fileExists := fileStatErr == nil

		if fileExists && skipExisting {
			continue
		} else if fileExists {
			// The skip existing flag was unset (false), so prompt user at run time
			fmt.Printf("This file already exists. Would you like to overwrite %s?\n", path.Base(file))
			fmt.Print("Enter Y to overwrite, or N to skip: ")
			reader := bufio.NewReader(os.Stdin)
			strOverwriteFile, _ := reader.ReadString('\n')
			if strings.TrimSpace(strOverwriteFile) != "Y" {
				continue
			}
		}

		u, _ := url.Parse(file)

		//Set initial jar with a cookie accepting the license agreement
		//Example taken from: https://gist.github.com/Rabbit52/a8a44c3c4cd514052952
		cookieJar.SetCookies(u, cookies)

		// Check to see if we are requested a file that requires authentication.
		// First, make sure we have access to the users username and password
		// and prompt if necessary.
		if !res.SkipAuth && requiresAuth {

			if len(otnPassword) == 0 {
				fmt.Printf("To complete the license acceptance, you must enter valid OTN credentials. Please enter your OTN password (%s):", otnUser)
				consolePass, _ := terminal.ReadPassword(int(syscall.Stdin))
				fmt.Println()
				otnPassword = string(consolePass)
			}

			fileResp = loginAndGetFile(file, client, otnUser, otnPassword)
			defer fileResp.Body.Close()

			requiresAuth = false
		} else {
			fileReq := getRequest(http.MethodGet, file, nil)
			fileResp = getResponse(fileReq, client)
			defer fileResp.Body.Close()
		}

		requestTotalSize := fileResp.ContentLength

		//Set up progress bar
		progressBar := pb.New64(requestTotalSize)
		progressBar.SetUnits(pb.U_BYTES)
		progressBar.Prefix(fmt.Sprintf("%s:", path.Base(file)))
		progressBar.Start()

		readerWithProgress := progressBar.NewProxyReader(fileResp.Body)
		savedFile, err := os.Create(path.Base(file))
		defer savedFile.Close()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(savedFile, readerWithProgress)
		if err != nil {
			log.Fatal(err)
		}
		progressBar.Finish()
	}

}
