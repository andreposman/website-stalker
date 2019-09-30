package terminal

import (
	"fmt"

	"github.com/andreposman/website-stalker/pkg/settings"
)

var (
	setting *settings.Setting
)

// Init ...
func init() {
	setting = settings.Build()
}

//Menu ...
func Menu() {
	fmt.Print("\n-----------------------------------")
	fmt.Print("| MENU |")
	fmt.Print("-----------------------------------\n\n")
	fmt.Println("1 - Stalk Websites 🕵")
	fmt.Println("2 - Show Logs 📄")
	fmt.Println("0 - Exit 🚪" + "\n")
	fmt.Print("------------------------------------------------------------------------------\n\n")
}

// Introduction ...
func Introduction() {

	name := setting.Name
	version := setting.Version

	fmt.Println(name + " - " + version)
}

// StalkingPrint ...
func StalkingPrint() {
	fmt.Print("\n-----------------------------------")
	fmt.Print("| STALKING...🕵 |")
	fmt.Print("-----------------------------\n\n")
}

// LogsPrint ...
func LogsPrint() {
	fmt.Print("\n-----------------------------------")
	fmt.Print("| LOGS 📄|")
	fmt.Print("-----------------------------------\n\n")
}

// ExitPrint ...
func ExitPrint() {
	fmt.Print("\n-----------------------------------")
	fmt.Print("| SEE YA LATER! 👋|")
	fmt.Print("-----------------------------------\n\n")
}
