package tester

import (
	"fmt"
	"net/http"
	"os"
	"website-stalker/src/logger"
)

// WebsiteTesting ...
func WebsiteTesting(url string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if response.StatusCode == 200 {
		fmt.Println("\n(✅) ➡ ", url, "was successfully loaded.")
		logger.RegisterLogs(url, true)
	} else {
		fmt.Println("(❌) ➡  Status Code:", response.StatusCode, "\nIt seems like the website", url, "is experiencing some problems.😞")
		logger.RegisterLogs(url, false)
	}
}
