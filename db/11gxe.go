package db

import "net/http"

type Resource struct {
	Component    string
	Version      string
	File         string
	License      string
	OS           string
	Arch         string
	AcceptCookie *http.Cookie
}

func GetXeResouces() []*Resource {

	acceptCookie := &http.Cookie{
		Name:   "oraclelicense",
		Value:  "accept-sqldev-cookie",
		Domain: ".oracle.com",
	}

	linuxXe := &Resource{
		Component:    "db",
		Version:      "11gXE",
		File:         "https://edelivery.oracle.com/akam/otn/linux/oracle11g/xe/oracle-xe-11.2.0-1.0.x86_64.rpm.zip",
		License:      "http://www.oracle.com/technetwork/licenses/database-11g-express-license-459621.html",
		OS:           "linux",
		Arch:         "amd64",
		AcceptCookie: acceptCookie,
	}

	windowsXe32 := &Resource{
		Component:    "db",
		Version:      "11gXE",
		File:         "https://edelivery.oracle.com/akam/otn/nt/oracle11g/xe/OracleXE112_Win32.zip",
		License:      "http://www.oracle.com/technetwork/licenses/database-11g-express-license-459621.html",
		OS:           "windows",
		Arch:         "32",
		AcceptCookie: acceptCookie,
	}

	windowsXe64 := &Resource{
		Component:    "db",
		Version:      "11gXE",
		File:         "https://edelivery.oracle.com/akam/otn/nt/oracle11g/xe/OracleXE112_Win64.zip",
		License:      "http://www.oracle.com/technetwork/licenses/database-11g-express-license-459621.html",
		OS:           "windows",
		Arch:         "amd64",
		AcceptCookie: acceptCookie,
	}

	xeResources := []*Resource{linuxXe, windowsXe32, windowsXe64}

	return xeResources
}
