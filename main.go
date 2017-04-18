package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tschf/odl/arch"
	"github.com/tschf/odl/dl"
	"github.com/tschf/odl/resource/finder"
)

// In order to download software, the OTN license agreement must be accepted.
// This is implemented by two mechanisms:
// 1. The command line argument `accept-license`
// 2. Prompting for user input, with a Y/N value where Y indicated truth
func getLicenseAcceptance(acceptFromFlag bool, licenseURL string, user string) bool {

	if acceptFromFlag {
		fmt.Printf("INFO: You (%s) have nominated you accept the OTN license agreement. Use of this software is tied to that acceptance.\n", user)
		return true
	}

	// Because the accept flag wasn't passed in, we want to prompt the user
	// to decide if they'd like to accept the license or not - passing in the
	// URL to the license agreement
	fmt.Println("Before continuing, you must accept the OTN license agreement.")
	fmt.Println(fmt.Sprintf("The full terms can be found here: %s", licenseURL))
	fmt.Print("Please enter Y if you accept the license agreement: ")

	reader := bufio.NewReader(os.Stdin)
	strLicenseAccepted, _ := reader.ReadString('\n')

	return strings.TrimSpace(strLicenseAccepted) == "Y"
}

func main() {

	var (
		flagUser          = flag.String("username", "", "Specify the user account that will be logging in and accepting the license agreement. Alternatively, set the environment variable OTN_USERNAME.")
		flagPassword      = flag.String("password", "", "Specify the password that corresponds to your OTN account. Alternatively, set the environment variable OTN_PASSWORD.")
		flagOs            = flag.String("os", "linux", "Specify the desired platform of the software. Should be \"linux\" or \"windows\"")
		flagComponent     = flag.String("component", "", "Specify the component to grab.")
		flagVersion       = flag.String("version", "", "Specify the software version. ")
		flagLang          = flag.String("lang", "na", "Specify the language of the software. Should be \"en\" or \"na\"")
		flagAcceptLicense = flag.Bool("accept-license", false, "Specify whether or not you accept the OTN license agreement for the nominated software.")
		flagArchitecture  arch.Arch
		flagSkipExisting  = flag.Bool("skip-if-exists", false, "Specify whether or not you want to skip existing files.")
	)

	flag.Var(&flagArchitecture, "arch", "Specify the desired architecture of the software. Should be \"x86\", \"x64\", or \"na\"")
	flag.Parse()

	otnUser := *flagUser
	if len(otnUser) == 0 {
		otnUser = os.Getenv("OTN_USERNAME")
	}

	if len(otnUser) == 0 {
		log.Fatal("You must specify an OTN username to access OTN files. Set with the flag -username or set the environment variable OTN_USERNAME.")
	}

	otnPassword := *flagPassword
	if len(otnPassword) == 0 {
		otnPassword = os.Getenv("OTN_PASSWORD")
	}

	selectedFile, ok := finder.FindResource(*flagComponent, *flagVersion, *flagOs, flagArchitecture, *flagLang)

	if ok {

		fmt.Printf("Beginning download process for %s %s\n", *flagComponent, *flagVersion)

		// The license accepted is done either through a command line flag
		// (accept-license), or if that is not set, prompted the user for input
		// which accepts a Y/N value. Only if the user inputs Y does that
		// indicate acceptance of the license.
		otnLicenseAccepted := getLicenseAcceptance(*flagAcceptLicense, selectedFile.License, otnUser)

		if !otnLicenseAccepted {
			fmt.Fprint(os.Stderr, "You must accept the license agreement in order to download. Exiting now.\n")
			os.Exit(1)
		}

		dl.SaveResource(selectedFile, otnUser, otnPassword, *flagSkipExisting)

		fmt.Println("Download complete.")
	} else {
		log.Fatal("Err, Could not find the selected file.")
	}
}
