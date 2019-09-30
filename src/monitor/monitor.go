package monitor

import (
	"fmt"
	"time"

	"github.com/andreposman/website-stalker/src/fileio"
	"github.com/andreposman/website-stalker/src/tester"
)

const quantityMonitoring = 3
const delay = 5

// StartMonitoring ...
func StartMonitoring() {
	websites := fileio.ReadWebsitesFromFile()

	for i := 0; i < quantityMonitoring; i++ {
		// o range dos meus sites podem me retornar duas coisas, o indice e o valor daquela determinada posição
		for i, url := range websites {
			fmt.Sprintf("Testing position %v of %v.\n", i, len(websites))
			fmt.Sprintf("Current website beeing tested is %v\n", url)
			tester.WebsiteTesting(url)
		}
		time.Sleep(delay * time.Second)
	}

	fmt.Print("\n-----------------------------------")
	fmt.Print("| Finished ✅ |")
	fmt.Print("-----------------------------------\n\n")
}
