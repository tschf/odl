package ords

import (
	"net/http"

	"github.com/tschf/odl/resource"
	"github.com/tschf/odl/types/arch"
)

func GetOrdsResources() []*resource.OracleResource {

	acceptCookie := &http.Cookie{
		Name:   "oraclelicense",
		Value:  "accept-sqldev-cookie",
		Domain: ".oracle.com",
	}

	ordsResources := []*resource.OracleResource{}

	ordsResources = append(ordsResources, &resource.OracleResource{
		Component:    "ords",
		Version:      "3.0.9",
		File:         "https://edelivery.oracle.com/akam/otn/java/ords/ords.3.0.9.348.07.16.zip",
		License:      "http://www.oracle.com/technetwork/licenses/sqldev-license-152021.html",
		OS:           "na",
		Arch:         arch.Na,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	ordsResources = append(ordsResources, &resource.OracleResource{
		Component:    "ords",
		Version:      "3.0.8",
		File:         "https://edelivery.oracle.com/akam/otn/java/ords/ords.3.0.8.277.08.01.zip",
		License:      "http://www.oracle.com/technetwork/licenses/sqldev-license-152021.html",
		OS:           "na",
		Arch:         arch.Na,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	ordsResources = append(ordsResources, &resource.OracleResource{
		Component:    "ords",
		Version:      "3.0.7",
		File:         "https://edelivery.oracle.com/akam/otn/java/ords/ords.3.0.7.253.09.40.zip",
		License:      "http://www.oracle.com/technetwork/licenses/sqldev-license-152021.html",
		OS:           "na",
		Arch:         arch.Na,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	ordsResources = append(ordsResources, &resource.OracleResource{
		Component:    "ords",
		Version:      "3.0.6",
		File:         "https://edelivery.oracle.com/akam/otn/java/ords/ords.3.0.6.176.08.46.zip",
		License:      "http://www.oracle.com/technetwork/licenses/sqldev-license-152021.html",
		OS:           "na",
		Arch:         arch.Na,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	ordsResources = append(ordsResources, &resource.OracleResource{
		Component:    "ords",
		Version:      "3.0.5",
		File:         "https://edelivery.oracle.com/akam/otn/java/ords/ords.3.0.5.124.10.54.zip",
		License:      "http://www.oracle.com/technetwork/licenses/sqldev-license-152021.html",
		OS:           "na",
		Arch:         arch.Na,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	ordsResources = append(ordsResources, &resource.OracleResource{
		Component:    "ords",
		Version:      "3.0.4",
		File:         "https://edelivery.oracle.com/akam/otn/java/ords/ords.3.0.4.60.12.48.zip",
		License:      "http://www.oracle.com/technetwork/licenses/sqldev-license-152021.html",
		OS:           "na",
		Arch:         arch.Na,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	ordsResources = append(ordsResources, &resource.OracleResource{
		Component:    "ords",
		Version:      "3.0.3",
		File:         "https://edelivery.oracle.com/akam/otn/java/ords/ords.3.0.3.351.13.24.zip",
		License:      "http://www.oracle.com/technetwork/licenses/sqldev-license-152021.html",
		OS:           "na",
		Arch:         arch.Na,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	ordsResources = append(ordsResources, &resource.OracleResource{
		Component:    "ords",
		Version:      "3.0.2",
		File:         "https://edelivery.oracle.com/akam/otn/java/ords/ords.3.0.2.294.08.40.zip",
		License:      "http://www.oracle.com/technetwork/licenses/sqldev-license-152021.html",
		OS:           "na",
		Arch:         arch.Na,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	ordsResources = append(ordsResources, &resource.OracleResource{
		Component:    "ords",
		Version:      "3.0.1",
		File:         "https://edelivery.oracle.com/akam/otn/java/ords/ords.3.0.1.177.18.02.zip",
		License:      "http://www.oracle.com/technetwork/licenses/sqldev-license-152021.html",
		OS:           "na",
		Arch:         arch.Na,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	ordsResources = append(ordsResources, &resource.OracleResource{
		Component:    "ords",
		Version:      "3.0",
		File:         "https://edelivery.oracle.com/akam/otn/java/ords/ords.3.0.0.121.10.23.zip",
		License:      "http://www.oracle.com/technetwork/licenses/sqldev-license-152021.html",
		OS:           "na",
		Arch:         arch.Na,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	return ordsResources
}
