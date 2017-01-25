package arch

import "fmt"

type Arch int

const (
	// Na Will be used to represent software that doesn't require a CPU
	// architecture. Making it the first entry in the list so that it becomes
	// the zero value. Useful for our arguments, so if it's omitted, it will
	// defualt to Na.
	Na Arch = iota
	X86
	X64
)

func (a *Arch) String() string {
	if *a == X86 {
		return "x86"
	} else if *a == X64 {
		return "x64"
	} else if *a == Na {
		return "na"
	} else {
		//Unspported value. Return an empty string
		return ""
	}
}

func (a *Arch) Set(v string) error {

	if v == "x86" {
		*a = X86
	} else if v == "x64" {
		*a = X64
	} else if v == "na" {
		*a = Na
	} else {
		return fmt.Errorf("Received value \"%s\" is an unknown architecture. Value must be \"x86\", \"x64\" or \"na\"", v)
	}

	return nil

}
