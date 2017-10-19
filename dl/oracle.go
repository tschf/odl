package dl

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

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

// The login form, which should point at: https://login.oracle.com/oaam_server/loginAuth.do
// Will have a 4 fields that get sent to the server in a POST request:
// 1. fk - a unique token that is returned in the form
// 2. clientOffset - The offset hours from GMT. e.g. Toronto would be -4
// 3. userid - the otn username
// 4. pass - the otn password
func submitLoginForm(oamLoginPageResp *http.Response, httpClient *http.Client, otnUser string, otnPassword string) *http.Response {
	goqueryDoc, err := goquery.NewDocumentFromResponse(oamLoginPageResp)
	if err != nil {
		log.Fatal(err)
	}
	loginForm := goqueryDoc.Find("form")

	loginFormData := url.Values{}
	loginAction, _ := loginForm.Attr("action")
	loginAction = loginURIPrefix + loginAction

	// 1. fk
	fk := loginForm.Find("input[name='fk']")
	fkVal, _ := fk.Attr("value")
	loginFormData.Set("fk", fkVal)

	// 2. clientOffset
	// The JavaScript function on the page looks like:
	// var time = new Date();
	// var offset = (time.getTimezoneOffset() / 60 * -1);
	t := time.Now()
	_, z := t.Zone()
	offset := strconv.Itoa(z / 60 / 60)
	loginFormData.Set("clientOffset", offset)

	// 3. userid
	loginFormData.Set("userid", otnUser)

	// 4. pass
	loginFormData.Set("pass", otnPassword)

	loginReq := getRequest(http.MethodPost, loginAction, loginFormData)
	loginResp := getResponse(loginReq, httpClient)

	return loginResp
}

// When login is required
func loginAndGetFile(file string, httpClient *http.Client, otnUser string, otnPassword string) *http.Response {

	fileRequest := getRequest(http.MethodGet, file, nil)

	resp := getResponse(fileRequest, httpClient)
	defer resp.Body.Close()

	// First response should contain a form that is expected to be posted to
	// https://login.oracle.com:443/oaam_server/oamLoginPage.jsp
	// It doesn't contain any login credentials, but contains a few bits of data
	// for the fields:
	// - TapSubmitURL
	// - tap_token
	// - OAM_REQ
	oamLoginPageResp := resubmitResponseForm(resp, httpClient)
	defer oamLoginPageResp.Body.Close()

	// Once that is returned, a login form is returned where we actually send
	// the login (username and password) data to.
	loginAuthResp := submitLoginForm(oamLoginPageResp, httpClient, otnUser, otnPassword)
	defer loginAuthResp.Body.Close()

	// Then we will submit our login data. Assuming authentication is valid
	// the next part is to send a GET request to:
	// https://login.oracle.com:443/oaam_server/authJump.do?jump=false
	// In the browser, this is achieve by way of a meta refresh tag. When requesting
	// this resource, it issues a 302 and return the results of the page at:
	// https://login.oracle.com:443/oaam_server/updateLoginStatus.do
	// Which in turns contains another form to submit.
	authJumpReq := getRequest(http.MethodGet, "https://login.oracle.com:443/oaam_server/authJump.do?jump=false", nil)
	authJumpResp := getResponse(authJumpReq, httpClient)
	defer authJumpResp.Body.Close()

	finalResp := resubmitResponseForm(authJumpResp, httpClient)

	return finalResp
}
