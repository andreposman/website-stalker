package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	showIntroduction()
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

func showIntroduction() {
	version := 1.1
	fmt.Println("Website Watcher - version ", version)
}

func showMenu() {
	fmt.Print("-----------------------------------\n\n")
	fmt.Println("1- Iniciar o Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func readUserInput() byte {
	var userInput byte
	fmt.Scan(&userInput)

	return userInput
}

func startMonitoring() {
	fmt.Println("Start watching...")
	url := "https://www.google.com"

	response, err := http.Get(url)

	fmt.Println(response.Status)

	if err != nil {
		fmt.Println(err)
	}

}
