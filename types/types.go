package types

import "net/http"

type Resource struct {
	Component    string
	Version      string
	File         string
	License      string
	OS           string
	Arch         string
	AcceptCookie *http.Cookie
}
