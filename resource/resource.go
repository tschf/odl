package resource

import (
	"net/http"

	"github.com/tschf/odl/arch"
)

// OracleResource represents the information about a resource that is hosted
// on the Oracle Tech Network (OTN) that can be downloaded.
type OracleResource struct {
	// Component is the high level detail about a particular resource, so that
	// resources can my logically grouped by their component. In general, the
	// the package name will be a 1-to-1 relationship with the component name.
	// However, some resources consist of a number of component names - such as
	// the instant client, which comes with sdk, basic, basic-lite, odbc, wrc
	// and sqlplus. This will be represented by the command argument "component"
	Component string
	// Version is used to determine which version of the particular component
	// to download.
	Version string
	// File represents the actual location of the specific software. This needs
	// to be a list, since some Oracle resources contain more than one file.
	// One such example is Oracle 12c EE. See:
	// http://www.oracle.com/technetwork/database/enterprise-edition/downloads/index.html
	File []string
	// License is a URL to the OTN license agreement to the specific software.
	// The license URL varies between each software download.
	License string
	// Java downloads, unlike most other Oracle resources does not seem to require
	// authentication. Instead, it requires only the accept cookie to be set.
	// Usually, the flow would be: request resource => go to login page =>
	// submit username/password (POST) and then receive the originally requested
	// file. In the case of Java downloads, the second step doesn't come into
	// play. So, we can use this flag (not the norm) to indicate when the
	// second step doesn't come into force
	SkipAuth bool
	// OS is for the operating system for the particular software. Valid values
	// would include: "linux", "osx" and "windows"
	OS string
	// Arch represents the architecture of the OS you would like to install the
	// software one. Values are provided through the arch package and include
	// x86, x64 or na (not applicable - as some software is architecture independent)
	Arch arch.Arch
	// Lange represents the language the software is targeting. In particular,
	// APEX comes with a version targeting just english. I think language targets
	// only happens with APEX. Valid values would be "en" or "na".
	Lang string
	// AcceptCookie is the http.Cookie that is required to indicate that the
	// License has been accepted. The cookie name has an every so slight variation
	// between the different software downloads, but usually starts with `accept-`
	// and ends in `-cookie`.
	AcceptCookie *http.Cookie
}
