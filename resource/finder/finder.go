package finder

import (
	"github.com/tschf/odl/arch"
	"github.com/tschf/odl/resource"
	"github.com/tschf/odl/resource/apex"
	"github.com/tschf/odl/resource/db"
	"github.com/tschf/odl/resource/instantclient"
	"github.com/tschf/odl/resource/java"
	"github.com/tschf/odl/resource/ords"
	"github.com/tschf/odl/resource/sqlcl"
	"github.com/tschf/odl/resource/sqldev"
)

func FindResource(componentName string, version string, os string, arch arch.Arch, lang string) (*resource.OracleResource, bool) {

	selectedComponent := resource.ComponentMap[componentName]

	var componentResources []*resource.OracleResource
	// first, load the relevant resources - target based on the component name
	// so we aren't uncessarily loading and looking at all resources across
	// components
	switch selectedComponent {
	case resource.APEX:
		componentResources = apex.GetApexResources()
	case resource.DB:
		componentResources = db.GetXeResouces()
		componentResources = append(componentResources, db.Get12cResouces()...)
	case resource.INSTANT_CLIENT:
		componentResources = instantclient.GetIcBasicLiteResources()
		componentResources = append(componentResources, instantclient.GetIcBasicResources()...)
		componentResources = append(componentResources, instantclient.GetIcBasicResources()...)
		componentResources = append(componentResources, instantclient.GetIcJdbcResources()...)
		componentResources = append(componentResources, instantclient.GetIcODBCResources()...)
		componentResources = append(componentResources, instantclient.GetIcSdkResources()...)
		componentResources = append(componentResources, instantclient.GetIcSqlplusResources()...)
		componentResources = append(componentResources, instantclient.GetIcWrcResources()...)
	case resource.JAVA:
		componentResources = java.GetJdkResources()
		componentResources = append(componentResources, java.GetJreResources()...)
	case resource.ORDS:
		componentResources = ords.GetOrdsResources()
	case resource.SQLCL:
		componentResources = sqlcl.GetSqlclResources()
	case resource.SQLDEV:
		componentResources = sqldev.GetSqldevResources()
	}

	// Interrogate all the components and look for the one matching all the
	// fields - language, architecture, version and OS.
	for _, el := range componentResources {
		if el.Component == componentName && el.Lang == lang && el.Version == version && el.OS == os && el.Arch == arch {
			return el, true
		}
	}

	// If we get to this point, there were no matching resources found, so we
	// should return an empty OracleResource, and false to indicate there
	// was no matching resource.
	return &resource.OracleResource{}, false

}
