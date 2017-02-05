package java

import (
	"net/http"

	"github.com/tschf/odl/resource"
	"github.com/tschf/odl/types/arch"
)

func GetJdk() []*resource.OracleResource {

	acceptCookie := &http.Cookie{
		Name:   "oraclelicense",
		Value:  "accept-securebackup-cookie",
		Domain: ".oracle.com",
	}

	jdkResources := []*resource.OracleResource{}

	jdkResources = append(jdkResources, &resource.OracleResource{
		Component:    "java-jdk",
		Version:      "8",
		File:         "https://edelivery.oracle.com/otn-pub/java/jdk/8u121-b13/e9e7ea248e2c4826b92b3f075a80e441/jdk-8u121-linux-x64.rpm",
		License:      "http://www.oracle.com/technetwork/java/javase/terms/license/index.html",
		SkipAuth:     true,
		OS:           "linux",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	jdkResources = append(jdkResources, &resource.OracleResource{
		Component:    "java-jdk",
		Version:      "8",
		File:         "https://edelivery.oracle.com/otn-pub/java/jdk/8u121-b13/e9e7ea248e2c4826b92b3f075a80e441/jdk-8u121-linux-x64.tar.gz",
		License:      "http://www.oracle.com/technetwork/java/javase/terms/license/index.html",
		SkipAuth:     true,
		OS:           "linux-other",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	jdkResources = append(jdkResources, &resource.OracleResource{
		Component:    "java-jdk",
		Version:      "8",
		File:         "https://edelivery.oracle.com/otn-pub/java/jdk/8u121-b13/e9e7ea248e2c4826b92b3f075a80e441/jdk-8u121-linux-i586.rpm",
		License:      "http://www.oracle.com/technetwork/java/javase/terms/license/index.html",
		SkipAuth:     true,
		OS:           "linux",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	jdkResources = append(jdkResources, &resource.OracleResource{
		Component:    "java-jdk",
		Version:      "8",
		File:         "https://edelivery.oracle.com/otn-pub/java/jdk/8u121-b13/e9e7ea248e2c4826b92b3f075a80e441/jdk-8u121-linux-i586.tar.gz",
		License:      "http://www.oracle.com/technetwork/java/javase/terms/license/index.html",
		SkipAuth:     true,
		OS:           "linux-other",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	jdkResources = append(jdkResources, &resource.OracleResource{
		Component:    "java-jdk",
		Version:      "8",
		File:         "https://edelivery.oracle.com/otn-pub/java/jdk/8u121-b13/e9e7ea248e2c4826b92b3f075a80e441/jdk-8u121-windows-x64.exe",
		License:      "http://www.oracle.com/technetwork/java/javase/terms/license/index.html",
		SkipAuth:     true,
		OS:           "windows",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	jdkResources = append(jdkResources, &resource.OracleResource{
		Component:    "java-jdk",
		Version:      "8",
		File:         "https://edelivery.oracle.com/otn-pub/java/jdk/8u121-b13/e9e7ea248e2c4826b92b3f075a80e441/jdk-8u121-windows-i586.exe",
		License:      "http://www.oracle.com/technetwork/java/javase/terms/license/index.html",
		SkipAuth:     true,
		OS:           "windows",
		Arch:         arch.X86,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	jdkResources = append(jdkResources, &resource.OracleResource{
		Component:    "java-jdk",
		Version:      "8",
		File:         "https://edelivery.oracle.com/otn-pub/java/jdk/8u121-b13/e9e7ea248e2c4826b92b3f075a80e441/jdk-8u121-macosx-x64.dmg",
		License:      "http://www.oracle.com/technetwork/java/javase/terms/license/index.html",
		SkipAuth:     true,
		OS:           "osx",
		Arch:         arch.X64,
		Lang:         "na",
		AcceptCookie: acceptCookie,
	})

	return jdkResources
}
