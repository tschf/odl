package sqlcl

import (
	"fmt"
	"log"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

// GetFilename returns the expect filename for the passed in, v, version of
// SQLcl
func GetFilename(v string) string {

	latestVersion := GetLatestVersion()

	return fmt.Sprintf("sqlcl-%s-no-jre.zip", latestVersion)

}

// GetLatestVersion returns the latest version of SQLcl by scraping the SQLcl
// download page. Returns purely the version number, e.g. 4.2.0.17.096.0933
func GetLatestVersion() string {

	doc, err := goquery.NewDocument("http://www.oracle.com/technetwork/developer-tools/sqlcl/downloads/index.html")
	if err != nil {
		log.Fatal(err)
	}

	// As at 8th April, 2016, the product version appears in a string on the page
	// resembling: April 6, 2017 - Update 4.2.0.17.096.0933
	//
	// <div class="sqldev-product">
	//     <h3 class="sqldev-header">
	//         <span class="sqldev-title">SQLcl 4.2</span>
	//     </h3>
	//     <p>
	//         <strong style="omitted">April 6, 2017 - Update 4.2.0.17.096.0933</strong>
	//     </p>
	//     <table style="width:100%">
	//     ..etc
	// </div>
	productInfo := doc.Find("div.sqldev-product > p")
	versionString := productInfo.Text()

	// \d{1,4} - match a sequence of digits ranging from 1 - 4 characters
	// (\.)? - optionally match a period at the end of the digit
	// +$ - repeat until we reach the end of the string/line
	versionRegexRule, _ := regexp.Compile(`(\d{1,4}(\.)?)+$`)
	extractedVersion := versionRegexRule.FindString(versionString)

	// return fmt.Sprintf("sqlcl-%s-no-jre.zip", extractedVersion)
	return extractedVersion
}
