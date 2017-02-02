package types

import (
	"net/http"

	"github.com/tschf/odl/types/arch"
)

type Resource struct {
	Component string
	Version   string
	File      string
	License   string
	// Java downloads, unlike most other Oracle resources does not seem to require
	// authentication. Instead, it requires only the accept cookie to be set.
	// Usually, the flow would be: request resource => go to login page =>
	// submit username/password (POST) and then receive the originally requested
	// file. In the case of Java downloads, the second step doesn't come into
	// play. So, we can use this flag (not the norm) to indicate when the
	// second step doesn't come into force
	SkipAuth     bool
	OS           string
	Arch         arch.Arch
	Lang         string
	AcceptCookie *http.Cookie
}
