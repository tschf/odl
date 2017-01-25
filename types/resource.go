package types

import (
	"net/http"

	"github.com/tschf/odl/types/arch"
)

type Resource struct {
	Component    string
	Version      string
	File         string
	License      string
	OS           string
	Arch         arch.Arch
	Lang         string
	AcceptCookie *http.Cookie
}
