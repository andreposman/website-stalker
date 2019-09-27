package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const quantityMonitoring = 3
const delay = 5

func main() {

	showIntroduction()

	for {
		showMenu()
		userInput := readUserInput()

		switch userInput {
		case 1:
			startMonitoring()
		case 2:
			fmt.Print("\n-----------------------------------")
			fmt.Print("| LOGS 📄|")
			fmt.Print("-----------------------------------\n\n")
			showLogs()
		case 0:
			fmt.Print("\n-----------------------------------")
			fmt.Print("| SEE YA LATER! 👋|")
			fmt.Print("-----------------------------------\n\n")
			os.Exit(0) //boa pratica sair com 0

		default:
			fmt.Println("Error")
			os.Exit(-1) // erro com o status -1

		}
	}

}

func showIntroduction() {
	version := 1.1
	fmt.Println("Website Watcher - version ", version)
}

func showMenu() {
	fmt.Print("\n-----------------------------------")
	fmt.Print("| MENU |")
	fmt.Print("-----------------------------------\n\n")
	fmt.Println("1 - Stalk Websites 🕵")
	fmt.Println("2 - Show Logs 📄")
	fmt.Println("0 - Exit 🚪" + "\n")
	fmt.Print("------------------------------------------------------------------------------\n\n")
}

func readUserInput() byte {
	var userInput byte
	fmt.Scan(&userInput)

	return userInput
}

func startMonitoring() {
	websites := readWebsitesFromFile()
	fmt.Print("\n-----------------------------------")
	fmt.Print("| STALKING...🕵 |")
	fmt.Print("-----------------------------\n\n")

	for i := 0; i < quantityMonitoring; i++ {
		// o range dos meus sites podem me retornar duas coisas, o indice e o valor daquela determinada posição
		for i, url := range websites {
			fmt.Sprintf("Testing position %v of %v.\n", i, len(websites))
			fmt.Sprintf("Current website beeing tested is %v\n", url)
			websiteTesting(url)
		}
		time.Sleep(delay * time.Second)
	}

	fmt.Print("\n-----------------------------------")
	fmt.Print("| Finished ✅ |")
	fmt.Print("-----------------------------------\n\n")
}

func websiteTesting(url string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if response.StatusCode == 200 {
		fmt.Println("\n(✅) ➡ ", url, "was successfully loaded.")
		registerLogs(url, true)
	} else {
		fmt.Println("(❌) ➡  Status Code:", response.StatusCode, "\nIt seems like the website", url, "is experiencing some problems.😞")
		registerLogs(url, false)
	}
}

func readWebsitesFromFile() []string {

	rawFile, err := os.Open("./websites.txt")
	var websites []string

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	file := bufio.NewReader(rawFile)

	for {
		line, err := file.ReadString('\n')
		line = strings.TrimSpace(line)
		websites = append(websites, line)

		if err == io.EOF {
			break
		} else if err != nil && err != io.EOF {
			fmt.Println(err)
			os.Exit(-1)
		}
	}

	rawFile.Close()
	return websites

}

func registerLogs(url string, status bool) {

	file, err := os.OpenFile("./logs/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " | " + url + " | online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func showLogs() {
	file, err := ioutil.ReadFile("./logs/log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
