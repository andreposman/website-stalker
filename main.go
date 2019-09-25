package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	showIntroduction()

	for {
		showMenu()
		userInput := readUserInput()

		switch userInput {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Logs")

		case 0:
			fmt.Println("Saindo do Programa")
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
	fmt.Println("-----------------------------------")
	fmt.Println("1- Iniciar o Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	fmt.Println("-----------------------------------")
	fmt.Println("")
}

func readUserInput() byte {
	var userInput byte
	fmt.Scan(&userInput)

	return userInput
}

func startMonitoring() {
	fmt.Println("\nStart watching...")

	websites := []string{"http://random-status-code.herokuapp.com/", "https://www.google.com", "http://andreposman.surge.sh/", "https://github.com"}

	// o range dos meus sites, podem me retornar duas coisas, o indice e o valor daquela determinada posição
	for i, url := range websites {
		fmt.Sprintf("Testing position %v of %v.\n", i, len(websites))
		fmt.Sprintf("Current website beeing tested is %v\n", url)
		websiteTesting(url)
	}

}

func websiteTesting(url string) {
	response, _ := http.Get(url)
	if response.StatusCode == 200 {
		fmt.Println("\n>> ", url, "was successfully loaded.")
	} else {
		fmt.Println("Status Code", response.StatusCode, "\nIt seems like the website", url, "is experiencing some problems.")
	}
}
