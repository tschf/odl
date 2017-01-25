package apex

import (
	"net/http"

	"github.com/tschf/odl/types"
	"github.com/tschf/odl/types/arch"
)

func GetApexResources() []*types.Resource {

	acceptCookie := &http.Cookie{
		Name:   "oraclelicense",
		Value:  "accept-appex3-cookie",
		Domain: ".oracle.com",
	}

	apex51 := &types.Resource{
		Component:    "apex",
		Version:      "5.1",
		File:         "https://edelivery.oracle.com/akam/otn/java/appexpress/apex_5.1.zip",
		License:      "http://www.oracle.com/technetwork/licenses/app-express-lic-152009.html",
		OS:           "na",
		Arch:         arch.Na,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	}

	apexResources := []*types.Resource{apex51}

	return apexResources
}
