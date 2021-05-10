package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const delay = 5
const monitoramentos = 5

func main() {
	//showNames()
	greetings()
	for  {
		showMenu()
		command := readCommand()
		executeCommand(command)
	}
}

func executeCommand(comando int) {
	switch comando {
	case 1:
		initMonitoring()
	case 2:
		fmt.Println("Logando aplicacao")
	case 0:
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
	fmt.Println("Monitoring applications...")
	// Arrays with fixed length
	sites := []string {"https://random-status-code.herokuapp.com/", "https://www.alura.com.br/", "https://www.caelum.com.br/"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Monitoring site ", i, ":", site)
			checkSite(site)
		}
		time.Sleep(delay * time.Second)
	}

}

func checkSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site", site, "foi carregado com problemas", "Status code", resp.StatusCode)
	}
}

func showNames(){
	names := []string {"Victor", "Livia", "Maria", "Joao"}
	names = append(names, "Felipe")
	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(cap(names)) // double the initial capacity
}