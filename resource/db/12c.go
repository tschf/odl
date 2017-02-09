package db

import (
	"net/http"

	"github.com/tschf/odl/arch"
	"github.com/tschf/odl/resource"
)

func Get12cResouces() []*resource.OracleResource {

	acceptCookie := &http.Cookie{
		Name:   "oraclelicense",
		Value:  "accept-dbindex-cookie",
		Domain: ".oracle.com",
	}

	db12cResources := []*resource.OracleResource{}

	// EE
	db12cResources = append(db12cResources, &resource.OracleResource{
		Component:    "db",
		Version:      "12.1.0.2.0EE",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/oracle12c/121020/linuxamd64_12102_database_1of2.zip", "https://edelivery.oracle.com/akam/otn/linux/oracle12c/121020/linuxamd64_12102_database_2of2.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/standard-license-152015.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	db12cResources = append(db12cResources, &resource.OracleResource{
		Component:    "db",
		Version:      "12.1.0.2.0EE",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/oracle12c/121020/winx64_12102_database_1of2.zip", "https://edelivery.oracle.com/akam/otn/nt/oracle12c/121020/winx64_12102_database_2of2.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/standard-license-152015.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	// SE2
	db12cResources = append(db12cResources, &resource.OracleResource{
		Component:    "db",
		Version:      "12.1.0.2.0SE2",
		File:         []string{"https://edelivery.oracle.com/akam/otn/linux/oracle12c/121020/linuxamd64_12102_database_se2_1of2.zip", "https://edelivery.oracle.com/akam/otn/linux/oracle12c/121020/linuxamd64_12102_database_se2_2of2.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/standard-license-152015.html",
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	db12cResources = append(db12cResources, &resource.OracleResource{
		Component:    "db",
		Version:      "12.1.0.2.0SE2",
		File:         []string{"https://edelivery.oracle.com/akam/otn/nt/oracle12c/121020/winx64_12102_SE2_database_1of2.zip", "https://edelivery.oracle.com/akam/otn/nt/oracle12c/121020/winx64_12102_SE2_database_2of2.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/standard-license-152015.html",
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	return db12cResources
}
