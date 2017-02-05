package instantclient

import (
	"net/http"

	"github.com/tschf/odl/arch"
	"github.com/tschf/odl/resource"
)

func GetIcSqlplusResources() []*resource.OracleResource {

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

	SqlplusResources := []*resource.OracleResource{}

	// 12.1.0.2.0

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/nt/instantclient/121020/instantclient-sqlplus-windows.x64-12.1.0.2.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx64,
	})

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/nt/instantclient/121020/instantclient-sqlplus-nt-12.1.0.2.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx86,
	})

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/oracle-instantclient12.1-sqlplus-12.1.0.2.0-1.x86_64.rpm",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/oracle-instantclient12.1-sqlplus-12.1.0.2.0-1.i386.rpm",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx86,
	})

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/instantclient-sqlplus-linux-12.1.0.2.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-other",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx86,
	})

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/instantclient-sqlplus-linux.x64-12.1.0.2.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-other",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/mac/instantclient/121020/instantclient-sqlplus-macos.x64-12.1.0.2.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/mac/instantclient/121020/instantclient-sqlplus-macos.x32-12.1.0.2.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	// 11.2.0.4.0

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/nt/instantclient/11204/instantclient-sqlplus-windows.x64-11.2.0.4.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx64,
	})

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/nt/instantclient/11204/instantclient-sqlplus-nt-11.2.0.4.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx86,
	})

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/oracle-instantclient11.2-sqlplus-11.2.0.4.0-1.x86_64.rpm",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/instantclient-sqlplus-linux.x64-11.2.0.4.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-other",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/oracle-instantclient11.2-sqlplus-11.2.0.4.0-1.i386.rpm",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/instantclient-sqlplus-linux-11.2.0.4.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-other",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/mac/instantclient/11204/instantclient-sqlplus-macos.x64-11.2.0.4.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/mac/instantclient/11204/instantclient-sqlplus-macos.x32-11.2.0.4.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	// 11.1.0.7.0

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "11.1.0.7.0",
		File:         "https://edelivery.oracle.com/akam/otn/nt/instantclient/111070/instantclient-sqlplus-win-x86-64-11.1.0.7.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx64,
	})

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "11.1.0.7.0",
		File:         "https://edelivery.oracle.com/akam/otn/nt/instantclient/111070/instantclient-sqlplus-win32-11.1.0.7.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx86,
	})

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "11.1.0.7.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/111070/oracle-instantclient11.1-sqlplus-11.1.0.7.0-1.x86_64.rpm",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "11.1.0.7.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/111070/sqlplus-11.1.0.7.0-linux-x86_64.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-other",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "11.1.0.7.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/111070/oracle-instantclient11.1-sqlplus-11.1.0.7.0-1.i386.rpm",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	SqlplusResources = append(SqlplusResources, &resource.OracleResource{
		Component:    "instantclient-sqlplus",
		Version:      "11.1.0.7.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/111070/instantclient-sqlplus-linux32-11.1.0.7.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-other",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	return SqlplusResources
}
