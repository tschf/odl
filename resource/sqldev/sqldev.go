package sqldev

import (
	"net/http"

	"github.com/tschf/odl/arch"
	"github.com/tschf/odl/resource"
)

func GetSqldevResources() []*resource.OracleResource {

	acceptCookie := &http.Cookie{
		Name:   "oraclelicense",
		Value:  "accept-sqldev-cookie",
		Domain: ".oracle.com",
	}

	sqldevResources := []*resource.OracleResource{}

	sqldevResources = append(sqldevResources, &resource.OracleResource{
		Component:    "sqldev",
		Version:      "4.1",
		File:         []string{"https://edelivery.oracle.com/akam/otn/java/sqldeveloper/sqldeveloper-4.1.5.21.78-1.noarch.rpm"},
		License:      "http://www.oracle.com/technetwork/licenses/sqldev-license-152021.html",
		OS:           "linux",
		Arch:         arch.Na,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	sqldevResources = append(sqldevResources, &resource.OracleResource{
		Component:    "sqldev",
		Version:      "4.1",
		File:         []string{"https://edelivery.oracle.com/akam/otn/java/sqldeveloper/sqldeveloper-4.1.5.21.78-no-jre.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/sqldev-license-152021.html",
		OS:           "other",
		Arch:         arch.Na,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	sqldevResources = append(sqldevResources, &resource.OracleResource{
		Component:    "sqldev",
		Version:      "4.1",
		File:         []string{"https://edelivery.oracle.com/akam/otn/java/sqldeveloper/sqldeveloper-4.1.5.21.78-macosx.app.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/sqldev-license-152021.html",
		OS:           "osx",
		Arch:         arch.Na,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	sqldevResources = append(sqldevResources, &resource.OracleResource{
		Component:    "sqldev",
		Version:      "4.1",
		File:         []string{"https://edelivery.oracle.com/akam/otn/java/sqldeveloper/sqldeveloper-4.1.5.21.78-no-jre.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/sqldev-license-152021.html",
		OS:           "windows",
		Arch:         arch.Na,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	sqldevResources = append(sqldevResources, &resource.OracleResource{
		Component:    "sqldev-jdk",
		Version:      "4.1",
		File:         []string{"https://edelivery.oracle.com/akam/otn/java/sqldeveloper/sqldeveloper-4.1.5.21.78-x64.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/sqldev-license-152021.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	return sqldevResources
}
