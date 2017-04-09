package sqlcl

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/tschf/odl/arch"
	"github.com/tschf/odl/resource"
)

func appendWithVersionCheck(resources []*resource.OracleResource, newResource *resource.OracleResource) []*resource.OracleResource {

	latestBuild := GetLatestVersion()

	// If the start of the latest build matches the major.minor version number
	// ensure the file URL points to that build
	if strings.HasPrefix(latestBuild, newResource.Version) {
		// We only expect one file path in the array, so just directly assign
		// the first value from the file list
		sqlclURL := newResource.File[0]

		fileBasename := filepath.Base(sqlclURL)
		latestBasename := GetFilename(latestBuild)
		sqlclNewURL := strings.Replace(sqlclURL, fileBasename, latestBasename, 1) //only expected to replace one occurrence

		newResource.File[0] = sqlclNewURL

	}

	resources = append(resources, newResource)

	return resources
}

func GetSqlclResources() []*resource.OracleResource {

	acceptCookie := &http.Cookie{
		Name:   "oraclelicense",
		Value:  "accept-sqldev-cookie",
		Domain: ".oracle.com",
	}

	sqlClResources := []*resource.OracleResource{}

	sqlClResources = appendWithVersionCheck(sqlClResources, &resource.OracleResource{
		Component:    "sqlcl",
		Version:      "4.2",
		File:         []string{"https://edelivery.oracle.com/akam/otn/java/sqldeveloper/sqlcl-4.2.0.17.096.0933-no-jre.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/sqldev-license-152021.html",
		OS:           "na",
		Arch:         arch.Na,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	return sqlClResources
}
