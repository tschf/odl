package instantclient

import (
	"net/http"

	"github.com/tschf/odl/arch"
	"github.com/tschf/odl/resource"
)

func GetIcPrecompilerResources() []*resource.OracleResource {

	acceptCookieIcWinx64 := &http.Cookie{
		Name:   "oraclelicense",
		Value:  "accept-ic_winx8664-cookie",
		Domain: ".oracle.com",
	}

	acceptCookieIcLinuxx64 := &http.Cookie{
		Name:   "oraclelicense",
		Value:  "accept-ic_linuxx8664-cookie",
		Domain: ".oracle.com",
	}

	acceptCookieIcLinuxx86 := &http.Cookie{
		Name:   "oraclelicense",
		Value:  "accept-ic_linux32-cookie",
		Domain: ".oracle.com",
	}

	acceptCookieIcWinx86 := &http.Cookie{
		Name:   "oraclelicense",
		Value:  "accept-ic_win32-cookie",
		Domain: ".oracle.com",
	}

	acceptCookieIcOSX := &http.Cookie{
		Name:   "oraclelicense",
		Value:  "accept-ic_solarisx8664-cookie",
		Domain: ".oracle.com",
	}

	PrecompilerResources := []*resource.OracleResource{}

	// 12.2.0.1.0

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "12.2.0.1.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/instantclient/122010/instantclient-precomp-windows.x64-12.2.0.1.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx64,
	})

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "12.2.0.1.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/instantclient/122010/instantclient-precomp-nt-12.2.0.1.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx86,
	})

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "12.2.0.1.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/122010/oracle-instantclient12.2-precomp-12.2.0.1.0-1.x86_64.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-rpm",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "12.2.0.1.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/122010/oracle-instantclient12.2-precomp-12.2.0.1.0-1.i386.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-rpm",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx86,
	})

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "12.2.0.1.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/122010/instantclient-precomp-linux-12.2.0.1.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx86,
	})

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "12.2.0.1.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/122010/instantclient-precomp-linux.x64-12.2.0.1.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	// 12.1.0.2.0

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/instantclient/121020/instantclient-precomp-windows.x64-12.1.0.2.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx64,
	})

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/instantclient/121020/instantclient-precomp-nt-12.1.0.2.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx86,
	})

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/oracle-instantclient12.1-precomp-12.1.0.2.0-1.x86_64.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-rpm",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/oracle-instantclient12.1-precomp-12.1.0.2.0-1.i386.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-rpm",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx86,
	})

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/instantclient-precomp-linux-12.1.0.2.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx86,
	})

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/instantclient-precomp-linux.x64-12.1.0.2.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/mac/instantclient/121020/instantclient-precomp-macos.x64-12.1.0.2.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/mac/instantclient/121020/instantclient-precomp-macos.x32-12.1.0.2.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	// 11.2.0.4.0

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/instantclient/11204/instantclient-precomp-windows.x64-11.2.0.4.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx64,
	})

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/instantclient/11204/instantclient-precomp-nt-11.2.0.4.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx86,
	})

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/oracle-instantclient11.2-basic-11.2.0.4.0-1.x86_64.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-rpm",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/instantclient-precomp-linux.x64-11.2.0.4.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/oracle-instantclient11.2-precomp-11.2.0.4.0-1.i386.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-rpm",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/instantclient-precomp-linux-11.2.0.4.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/mac/instantclient/11204/instantclient-precomp-macos.x64-11.2.0.4.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	PrecompilerResources = append(PrecompilerResources, &resource.OracleResource{
		Component:    "instantclient-precompiler",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/mac/instantclient/11204/instantclient-precomp-macos.x32-11.2.0.4.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	return PrecompilerResources
}
