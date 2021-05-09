package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	greetings()
	showMenu()
	command := readCommand()
	executeCommand(command)
}

func executeCommand(comando int) {
	switch comando {
	case 1:
		initMonitoring()
	case 2:
		fmt.Println("Logando aplicacao")
	case 3:
		fmt.Println("Saindo da aplicacao")
		os.Exit(0)
	default:
		fmt.Println("Nao conheco esse comando")
	}
}

func showMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do Programa")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("Comando escolhido:", command)
	return command
}

func greetings() {
	fmt.Println("Ola Mr. ? What is your name ?")
	name, _ := getNameAndAge()
	version := 1.1
	fmt.Println("Hello Mr.", name)
	fmt.Println("Program version ", version)
}

func getNameAndAge() (string, int) {
	name := "Victor Alcantara"
	age := 31
	return name, age
}

func initMonitoring() {
	fmt.Println("Monitorando aplicacao")
	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)
	fmt.Println(resp)
}