package sqlcl

import (
	"net/http"

	"github.com/tschf/odl/arch"
	"github.com/tschf/odl/resource"
)

func GetSqlclResources() []*resource.OracleResource {

	acceptCookie := &http.Cookie{
		Name:   "oraclelicense",
		Value:  "accept-sqldev-cookie",
		Domain: ".oracle.com",
	}

	sqlClResources := []*resource.OracleResource{}

	sqlClResources = append(sqlClResources, &resource.OracleResource{
		Component:    "sqlcl",
		Version:      "4.2",
		File:         "https://edelivery.oracle.com/akam/otn/java/sqldeveloper/sqlcl-4.2.0.16.355.0402-no-jre.zip",
		License:      "http://www.oracle.com/technetwork/licenses/sqldev-license-152021.html",
		OS:           "na",
		Arch:         arch.Na,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	return sqlClResources
}
