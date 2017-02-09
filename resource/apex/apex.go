package apex

import (
	"net/http"

	"github.com/tschf/odl/arch"
	"github.com/tschf/odl/resource"
)

func GetApexResources() []*resource.OracleResource {

	acceptCookie := &http.Cookie{
		Name:   "oraclelicense",
		Value:  "accept-appex3-cookie",
		Domain: ".oracle.com",
	}

	apexResources := []*resource.OracleResource{}

	// Add APEX 5.1 - the version for all languages
	apexResources = append(apexResources, &resource.OracleResource{
		Component:    "apex",
		Version:      "5.1",
		File:         []string{"https://edelivery.oracle.com/akam/otn/java/appexpress/apex_5.1.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/app-express-lic-152009.html",
		OS:           "na",
		Arch:         arch.Na,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	// Add APEX 5.1 - the version for english
	apexResources = append(apexResources, &resource.OracleResource{
		Component:    "apex",
		Version:      "5.1",
		File:         []string{"https://edelivery.oracle.com/akam/otn/java/appexpress/apex_5.1_en.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/app-express-lic-152009.html",
		OS:           "na",
		Arch:         arch.Na,
		Lang:         "en",
		AcceptCookie: acceptCookie,
	})

	// Add APEX 5.0 - the version for all languages
	apexResources = append(apexResources, &resource.OracleResource{
		Component:    "apex",
		Version:      "5.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/java/appexpress/apex_5.0.4.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/app-express-lic-152009.html",
		OS:           "na",
		Arch:         arch.Na,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	// Add APEX 5.0 - the version for english
	apexResources = append(apexResources, &resource.OracleResource{
		Component:    "apex",
		Version:      "5.0",
		File:         []string{"https://edelivery.oracle.com/akam/otn/java/appexpress/apex_5.0.4_en.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/app-express-lic-152009.html",
		OS:           "na",
		Arch:         arch.Na,
		Lang:         "en",
		AcceptCookie: acceptCookie,
	})

	// Add APEX 4.2 - the version for all languages
	apexResources = append(apexResources, &resource.OracleResource{
		Component:    "apex",
		Version:      "4.2",
		File:         []string{"https://edelivery.oracle.com/akam/otn/java/appexpress/apex_4.2.6.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/app-express-lic-152009.html",
		OS:           "na",
		Arch:         arch.Na,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	// Add APEX 4.2 - the version for english
	apexResources = append(apexResources, &resource.OracleResource{
		Component:    "apex",
		Version:      "4.2",
		File:         []string{"https://edelivery.oracle.com/akam/otn/java/appexpress/apex_4.2.6_en.zip"},
		License:      "http://www.oracle.com/technetwork/licenses/app-express-lic-152009.html",
		OS:           "na",
		Arch:         arch.Na,
		Lang:         "en",
		AcceptCookie: acceptCookie,
	})

	return apexResources
}
