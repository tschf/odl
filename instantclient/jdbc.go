package instantclient

import (
	"net/http"

	"github.com/tschf/odl/types"
	"github.com/tschf/odl/types/arch"
)

func GetIcJdbcResources() []*types.Resource {

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

	JdbcResources := []*types.Resource{}

	// 12.1.0.2.0

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/nt/instantclient/121020/instantclient-jdbc-windows.x64-12.1.0.2.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx64,
	})

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/nt/instantclient/121020/instantclient-jdbc-nt-12.1.0.2.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx86,
	})

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/oracle-instantclient12.1-jdbc-12.1.0.2.0-1.x86_64.rpm",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/oracle-instantclient12.1-jdbc-12.1.0.2.0-1.i386.rpm",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx86,
	})

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/instantclient-jdbc-linux-12.1.0.2.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-other",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx86,
	})

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/121020/instantclient-jdbc-linux.x64-12.1.0.2.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-other",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/mac/instantclient/121020/instantclient-jdbc-macos.x64-12.1.0.2.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "12.1.0.2.0",
		File:         "https://edelivery.oracle.com/akam/otn/mac/instantclient/121020/instantclient-jdbc-macos.x32-12.1.0.2.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	// 11.2.0.4.0

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/nt/instantclient/11204/instantclient-jdbc-windows.x64-11.2.0.4.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx64,
	})

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/nt/instantclient/11204/instantclient-jdbc-nt-11.2.0.4.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx86,
	})

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/oracle-instantclient11.2-jdbc-11.2.0.4.0-1.x86_64.rpm",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/instantclient-jdbc-linux.x64-11.2.0.4.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-other",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/oracle-instantclient11.2-jdbc-11.2.0.4.0-1.i386.rpm",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/11204/instantclient-jdbc-linux-11.2.0.4.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-other",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/mac/instantclient/11204/instantclient-jdbc-macos.x64-11.2.0.4.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "11.2.0.4.0",
		File:         "https://edelivery.oracle.com/akam/otn/mac/instantclient/11204/instantclient-jdbc-macos.x32-11.2.0.4.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "osx",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcOSX,
	})

	// 11.1.0.7.0

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "11.1.0.7.0",
		File:         "https://edelivery.oracle.com/akam/otn/nt/instantclient/111070/instantclient-jdbc-win-x86-64-11.1.0.7.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx64,
	})

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "11.1.0.7.0",
		File:         "https://edelivery.oracle.com/akam/otn/nt/instantclient/111070/instantclient-jdbc-win32-11.1.0.7.0.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcWinx86,
	})

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "11.1.0.7.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/111070/oracle-instantclient11.1-jdbc-11.1.0.7.0-1.x86_64.rpm",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "11.1.0.7.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/111070/jdbc-11.1.0.7.0-linux-x86_64.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-other",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "11.1.0.7.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/111070/oracle-instantclient11.1-jdbc-11.1.0.7.0-1.i386.rpm",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	JdbcResources = append(JdbcResources, &types.Resource{
		Component:    "instantclient-jdbc",
		Version:      "11.1.0.7.0",
		File:         "https://edelivery.oracle.com/akam/otn/linux/instantclient/111070/instantclient-jdbc-linux32-11.1.0.7.zip",
		License:      "http://www.oracle.com/technetwork/licenses/distribution-license-152002.html",
		OS:           "linux-other",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookieIcLinuxx64,
	})

	return JdbcResources
}
