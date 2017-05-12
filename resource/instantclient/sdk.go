package instantclient

import (
	"net/http"

	"github.com/tschf/odl/arch"
	"github.com/tschf/odl/resource"
)

func GetIcSdkResources() []*resource.OracleResource {

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

	SdkResources := []*resource.OracleResource{}

	// 12.1.0.2.0

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/instantclient/121020/instantclient-sdk-windows.x64-12.1.0.2.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx64,
	})

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/instantclient/121020/instantclient-sdk-nt-12.1.0.2.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx86,
	})

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/oracle-instantclient12.1-devel-12.1.0.2.0-1.x86_64.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-rpm",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/oracle-instantclient12.1-devel-12.1.0.2.0-1.i386.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-rpm",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx86,
	})

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/instantclient-sdk-linux-12.1.0.2.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx86,
	})

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/instantclient-sdk-linux.x64-12.1.0.2.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/mac/instantclient/121020/instantclient-sdk-macos.x64-12.1.0.2.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "12.1.0.2.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/mac/instantclient/121020/instantclient-sdk-macos.x32-12.1.0.2.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	// 11.2.0.4.0

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/instantclient/11204/instantclient-sdk-windows.x64-11.2.0.4.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx64,
	})

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/instantclient/11204/instantclient-sdk-nt-11.2.0.4.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx86,
	})

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/oracle-instantclient11.2-devel-11.2.0.4.0-1.x86_64.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-rpm",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/instantclient-sdk-linux.x64-11.2.0.4.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/oracle-instantclient11.2-devel-11.2.0.4.0-1.i386.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-rpm",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/instantclient-sdk-linux-11.2.0.4.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/mac/instantclient/11204/instantclient-sdk-macos.x64-11.2.0.4.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "11.2.0.4.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/mac/instantclient/11204/instantclient-sdk-macos.x32-11.2.0.4.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	// 11.1.0.7.0

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "11.1.0.7.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/instantclient/111070/instantclient-sdk-win-x86-64-11.1.0.7.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx64,
	})

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "11.1.0.7.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/instantclient/111070/instantclient-sdk-win32-11.1.0.7.0.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx86,
	})

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "11.1.0.7.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/111070/oracle-instantclient11.1-devel-11.1.0.7.0-1.x86_64.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-rpm",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "11.1.0.7.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/111070/sdk-11.1.0.7.0-linux-x86_64.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "11.1.0.7.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/111070/oracle-instantclient11.1-devel-11.1.0.7.0-1.i386.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-rpm",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	SdkResources = append(SdkResources, &resource.OracleResource{
		Component:    "instantclient-sdk",
		Version:      "11.1.0.7.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/instantclient/111070/instantclient-sdk-linux32-11.1.0.7.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	return SdkResources
}
