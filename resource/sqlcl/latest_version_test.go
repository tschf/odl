package sqlcl

import (
	"strings"
	"testing"
)

func TestLatestVersion(t *testing.T) {

	t.Log("Testing sqlcl version")

	latestVersion := GetLatestVersion()

	if len(latestVersion) == 0 {
		t.Errorf("Version not parsed in result")
	}

	buildCmps := strings.Split(latestVersion, ".")
	const expectedCmpsLen = 6 //4.2.0.17.096.0933

	if len(buildCmps) != expectedCmpsLen {
		t.Errorf("Build version (%s) not as expected. Should having 6 components similar to 4.2.0.17.096.0933", latestVersion)
	}

}
