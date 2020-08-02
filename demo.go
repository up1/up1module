package demo

import "github.com/fatih/color"

// SayHi : dddddd
func SayHi(name string) string {
	if name == "" {
		return "Error"
	}
	return say(name)
}

func say(name string) string {
	color.Cyan("Prints text in cyan.")
	return "Say hi " + name
}
