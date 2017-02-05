package instantclient

import (
	"net/http"

	"github.com/tschf/odl/resource"
	"github.com/tschf/odl/types/arch"
)

func GetIcODBCResources() []*resource.OracleResource {

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

	OdbcResources := []*resource.OracleResource{}

	// 12.1.0.2.0

	OdbcResources = append(OdbcResources, &resource.OracleResource{
		Component:    "instantclient-odbc",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/nt/instantclient/121020/instantclient-odbc-windows.x64-12.1.0.2.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx64,
	})

	OdbcResources = append(OdbcResources, &resource.OracleResource{
		Component:    "instantclient-odbc",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/nt/instantclient/121020/instantclient-odbc-nt-12.1.0.2.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx86,
	})

	OdbcResources = append(OdbcResources, &resource.OracleResource{
		Component:    "instantclient-odbc",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/oracle-instantclient12.1-odbc-12.1.0.2.0-1.x86_64.rpm",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	OdbcResources = append(OdbcResources, &resource.OracleResource{
		Component:    "instantclient-odbc",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/oracle-instantclient12.1-odbc-12.1.0.2.0-1.i386.rpm",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx86,
	})

	OdbcResources = append(OdbcResources, &resource.OracleResource{
		Component:    "instantclient-odbc",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/instantclient-odbc-linux-12.1.0.2.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-other",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx86,
	})

	OdbcResources = append(OdbcResources, &resource.OracleResource{
		Component:    "instantclient-odbc",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/instantclient-odbc-linux.x64-12.1.0.2.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-other",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	OdbcResources = append(OdbcResources, &resource.OracleResource{
		Component:    "instantclient-odbc",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/mac/instantclient/121020/instantclient-odbc-macos.x64-12.1.0.2.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	OdbcResources = append(OdbcResources, &resource.OracleResource{
		Component:    "instantclient-odbc",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/nt/instantclient/121020/instantclient-odbc-nt-12.1.0.2.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx86,
	})

	OdbcResources = append(OdbcResources, &resource.OracleResource{
		Component:    "instantclient-odbc",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/mac/instantclient/121020/instantclient-odbc-macos.x32-12.1.0.2.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	// 11.2.0.4.0

	OdbcResources = append(OdbcResources, &resource.OracleResource{
		Component:    "instantclient-odbc",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/nt/instantclient/11204/instantclient-odbc-windows.x64-11.2.0.4.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx64,
	})

	OdbcResources = append(OdbcResources, &resource.OracleResource{
		Component:    "instantclient-odbc",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/oracle-instantclient11.2-odbc-11.2.0.4.0-1.x86_64.rpm",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	OdbcResources = append(OdbcResources, &resource.OracleResource{
		Component:    "instantclient-odbc",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/instantclient-odbc-linux.x64-11.2.0.4.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-other",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	OdbcResources = append(OdbcResources, &resource.OracleResource{
		Component:    "instantclient-odbc",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/oracle-instantclient11.2-odbc-11.2.0.4.0-1.i386.rpm",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	OdbcResources = append(OdbcResources, &resource.OracleResource{
		Component:    "instantclient-odbc",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/instantclient-odbc-linux-11.2.0.4.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-other",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	OdbcResources = append(OdbcResources, &resource.OracleResource{
		Component:    "instantclient-odbc",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/nt/instantclient/11204/instantclient-odbc-nt-11.2.0.4.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx86,
	})

	// 11.1.0.7.0

	OdbcResources = append(OdbcResources, &resource.OracleResource{
		Component:    "instantclient-odbc",
		Version:      "11.1.0.7.0",
		File:         "https://edelivery.oracle.com/akam/otn/nt/instantclient/111070/instantclient-odbc-win-x86-64-11.1.0.7.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx64,
	})

	OdbcResources = append(OdbcResources, &resource.OracleResource{
		Component:    "instantclient-odbc",
		Version:      "11.1.0.7.0",
		File:         "https://edelivery.oracle.com/akam/otn/nt/instantclient/111070/instantclient-odbc-win32-11.1.0.7.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx86,
	})

	OdbcResources = append(OdbcResources, &resource.OracleResource{
		Component:    "instantclient-odbc",
		Version:      "11.1.0.7.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/111070/oracle-instantclient11.1-odbc-11.1.0.7.0-1.i386.rpm",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	OdbcResources = append(OdbcResources, &resource.OracleResource{
		Component:    "instantclient-odbc",
		Version:      "11.1.0.7.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/111070/instantclient-odbc-linux32-11.1.0.7.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-other",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	return OdbcResources
}
