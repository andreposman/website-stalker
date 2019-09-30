package main

import (
	"fmt"
	"os"

	"github.com/andreposman/website-stalker/src/logger"
	"github.com/andreposman/website-stalker/src/monitor"
	"github.com/andreposman/website-stalker/src/terminal"
)

func main() {
	terminal.Introduction()

	for {
		terminal.Menu()
		userInput := readUserInput()

		switch userInput {
		case 1:
			monitor.StartMonitoring()
		case 2:
			logger.ShowLogs()
		case 0:
			terminal.ExitPrint()
			os.Exit(0)

		default:
			fmt.Println("Error")
			os.Exit(-1)

		}
	}

}

func readUserInput() byte {
	var userInput byte
	fmt.Scan(&userInput)

	return userInput
}
