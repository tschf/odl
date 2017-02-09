package resource

type ComponentType int

// ComponentPackageMap is an alias to a map[string]string. It is used to store
// component names to relate to which package they belong to. For example
// the instantclient package consists of: instantclient-sqlplus, instantclient-basic
// etc. So I want a reference to know which component belongs to which package
// to aid in determining which package to pull from
type ComponentPackageMap map[string]ComponentType

const (
	APEX ComponentType = iota
	DB
	INSTANT_CLIENT
	JAVA
	ORDS
	SQLCL
	SQLDEV
)

var (
	ComponentMap = ComponentPackageMap{
		"apex": APEX,
		"db":   DB,
		"instantclient-basic":      INSTANT_CLIENT,
		"instantclient-basic-lite": INSTANT_CLIENT,
		"instantclient-jdbc":       INSTANT_CLIENT,
		"instantclient-odbc":       INSTANT_CLIENT,
		"instantclient-sdk":        INSTANT_CLIENT,
		"instantclient-sqlplus":    INSTANT_CLIENT,
		"instantclient-wrc":        INSTANT_CLIENT,
		"java-jdk":                 JAVA,
		"java-jre":                 JAVA,
		"ords":                     ORDS,
		"sqlcl":                    SQLCL,
		"sqldev":                   SQLDEV,
	}
)
