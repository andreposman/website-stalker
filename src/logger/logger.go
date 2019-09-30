package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/andreposman/website-stalker/src/terminal"
)

// RegisterLogs ...
func RegisterLogs(url string, status bool) {
	terminal.LogsPrint()
	file, err := os.OpenFile("../logs/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " | " + url + " | online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

// ShowLogs ...
func ShowLogs() {
	file, err := ioutil.ReadFile("../logs/log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
