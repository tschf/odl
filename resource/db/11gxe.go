package db

import (
	"net/http"

	"github.com/tschf/odl/arch"
	"github.com/tschf/odl/resource"
)

func GetXeResouces() []*resource.OracleResource {

	acceptCookie := &http.Cookie{
		Name:   "oraclelicense",
		Value:  "accept-sqldev-cookie",
		Domain: ".oracle.com",
	}

	xeResources := []*resource.OracleResource{}

	// Oracle 11gXE for Linux, 64-bit
	xeResources = append(xeResources, &resource.OracleResource{
		Component:    "db",
		Version:      "11gXE",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/oracle11g/xe/oracle-xe-11.2.0-1.0.x86_64.rpm.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/database-11g-express-license-459621.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	// Oracle 11gXE for Windows, 32-bit
	xeResources = append(xeResources, &resource.OracleResource{
		Component:    "db",
		Version:      "11gXE",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/oracle11g/xe/OracleXE112_Win32.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/database-11g-express-license-459621.html",
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	// Oracle 11gXE for windows, 64bit
	xeResources = append(xeResources, &resource.OracleResource{
		Component:    "db",
		Version:      "11gXE",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/oracle11g/xe/OracleXE112_Win64.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/database-11g-express-license-459621.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	return xeResources
}
