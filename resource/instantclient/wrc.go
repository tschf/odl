package instantclient

import (
	"net/http"

	"github.com/tschf/odl/arch"
	"github.com/tschf/odl/resource"
)

func GetIcWrcResources() []*resource.OracleResource {

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

	WrcResources := []*resource.OracleResource{}

	// 12.2.0.1.0

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "12.2.0.1.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/instantclient/122010/instantclient-tools-windows.x64-12.2.0.1.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx64,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "12.2.0.1.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/instantclient/122010/instantclient-tools-nt-12.2.0.1.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx86,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "12.2.0.1.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/122010/oracle-instantclient12.2-tools-12.2.0.1.0-1.x86_64.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-rpm",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "12.2.0.1.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/122010/oracle-instantclient12.2-tools-12.2.0.1.0-1.i386.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-rpm",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx86,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "12.2.0.1.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/122010/instantclient-tools-linux-12.2.0.1.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx86,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "12.2.0.1.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/122010/instantclient-tools-linux.x64-12.2.0.1.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	// WrcResources = append(WrcResources, &resource.OracleResource{
	//     Component:    "instantclient-wrc",
	//     Version:      "12.2.0.1.0",
	//     File:         []string{"https://edelivery.oracle.com/akam/otn/mac/instantclient/122010/instantclient-tools-macos.x64-12.2.0.1.0.zip"},
	//     License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
	//     OS:           "osx",
	//     Arch:         arch.X64,
	//     Lang:         "na",
	//     AcceptCookie: acceptCookieIcOSX,
	// })
	//
	// WrcResources = append(WrcResources, &resource.OracleResource{
	//     Component:    "instantclient-wrc",
	//     Version:      "12.2.0.1.0",
	//     File:         []string{"https://edelivery.oracle.com/akam/otn/mac/instantclient/122010/instantclient-tools-macos.x32-12.2.0.1.0.zip"},
	//     License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
	//     OS:           "osx",
	//     Arch:         arch.X86,
	//     Lang:         "na",
	//     AcceptCookie: acceptCookieIcOSX,
	// })

	// 12.1.0.2.0

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/instantclient/121020/instantclient-tools-windows.x64-12.1.0.2.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx64,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/instantclient/121020/instantclient-tools-nt-12.1.0.2.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx86,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/oracle-instantclient12.1-tools-12.1.0.2.0-1.x86_64.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-rpm",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/oracle-instantclient12.1-tools-12.1.0.2.0-1.i386.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-rpm",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx86,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/instantclient-tools-linux-12.1.0.2.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx86,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/instantclient-tools-linux.x64-12.1.0.2.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/mac/instantclient/121020/instantclient-tools-macos.x64-12.1.0.2.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/mac/instantclient/121020/instantclient-tools-macos.x32-12.1.0.2.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	// 11.2.0.4.0

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/instantclient/11204/instantclient-tools-windows.x64-11.2.0.4.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx64,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/instantclient/11204/instantclient-tools-nt-11.2.0.4.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx86,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/oracle-instantclient11.2-tools-11.2.0.4.0-1.x86_64.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-rpm",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/instantclient-tools-linux.x64-11.2.0.4.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/oracle-instantclient11.2-tools-11.2.0.4.0-1.i386.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-rpm",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/instantclient-tools-linux-11.2.0.4.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/mac/instantclient/11204/instantclient-tools-macos.x64-11.2.0.4.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/mac/instantclient/11204/instantclient-tools-macos.x32-11.2.0.4.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	// 11.1.0.7.0

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "11.1.0.7.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/instantclient/111070/instantclient-tools-win-x86-64-11.1.0.7.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx64,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "11.1.0.7.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/instantclient/111070/instantclient-tools-win32-11.1.0.7.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx86,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "11.1.0.7.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/111070/oracle-instantclient11.1-tools-11.1.0.7.0-1.x86_64.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-rpm",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "11.1.0.7.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/111070/tools-11.1.0.7.0-linux-x86_64.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "11.1.0.7.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/111070/oracle-instantclient11.1-tools-11.1.0.7.0-1.i386.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-rpm",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})
	//
	WrcResources = append(WrcResources, &resource.OracleResource{
		Component:    "instantclient-wrc",
		Version:      "11.1.0.7.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/111070/instantclient-tools-linux32-11.1.0.7.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	return WrcResources
}
