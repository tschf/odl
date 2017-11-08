package dl

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const loginURIPrefix string = "https://login.oracle.com"

func resubmitResponseForm(resp *http.Response, httpClient *http.Client) *http.Response {
	goqueryDoc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}
	pageForm := goqueryDoc.Find("form")

	pageAction, _ := pageForm.Attr("action")
	if !strings.HasPrefix(pageAction, loginURIPrefix) {
		pageAction = loginURIPrefix + pageAction
	}

	formInputData := url.Values{}
	pageForm.Find("input").Each(func(index int, el *goquery.Selection) {
		inputName, _ := el.Attr("name")
		inputValue, _ := el.Attr("value")

		if len(inputName) > 0 {
			formInputData.Set(inputName, inputValue)
		}
	})

	formReq := getRequest(http.MethodPost, pageAction, formInputData)
	formResp := getResponse(formReq, httpClient)

	return formResp
}

// The login form, which should point at: /oam/server/sso/auth_cred_submit
// Will have a 6 fields that get sent to the server in a POST request:
// - v
// - OAM_REQ
// - site2pstoretoken
// - locale
// - ssousername
// - password
func submitLoginForm(oamLoginPageResp *http.Response, httpClient *http.Client, otnUser string, otnPassword string) *http.Response {
	goqueryDoc, err := goquery.NewDocumentFromResponse(oamLoginPageResp)
	if err != nil {
		log.Fatal(err)
	}

	loginForm := goqueryDoc.Find("form")

	loginFormData := url.Values{}
	loginAction, _ := loginForm.Attr("action")
	loginAction = loginURIPrefix + loginAction

	loginForm.Find("input").Each(func(index int, el *goquery.Selection) {
		inputName, _ := el.Attr("name")
		inputValue, _ := el.Attr("value")

		if len(inputName) > 0 {
			loginFormData.Set(inputName, inputValue)
		}
	})

	// The login form has fields in the input element names:
	// - ssousername
	// - password
	loginFormData.Set("ssousername", otnUser)
	loginFormData.Set("password", otnPassword)

	loginReq := getRequest(http.MethodPost, loginAction, loginFormData)
	loginResp := getResponse(loginReq, httpClient)

	return loginResp
}

// When login is required
func loginAndGetFile(file string, httpClient *http.Client, otnUser string, otnPassword string) *http.Response {

	// Request the file, e.g. https://edelivery.oracle.com/akam/otn/java/appexpress/apex_5.1.zip
	// This will result in HTTP 302 to: /oam/server/osso_login?..
	// This page contains a form, with the action to: /mysso/signon.jsp
	fileRequest := getRequest(http.MethodGet, file, nil)
	resp := getResponse(fileRequest, httpClient)
	defer resp.Body.Close()

	// Here, we post the form: /mysso/signon.jsp with all the input elements.
	// The response contains another form
	oamLoginPageResp := resubmitResponseForm(resp, httpClient)
	defer oamLoginPageResp.Body.Close()

	// Once that is returned, a login form is returned where we actually send
	// the login (username and password) data to.
	// This is the last piece that will have a redirect to download the file.
	// (at least, until things change again)
	loginAuthResp := submitLoginForm(oamLoginPageResp, httpClient, otnUser, otnPassword)

	return loginAuthResp
}
