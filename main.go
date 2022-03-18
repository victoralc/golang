package main

import (
	"fmt"
)

func main() {
	nome := "Victor Alcantara"
	versao := 1.1
	idade := 31
	fmt.Println("Ola, sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa esta na versao,", versao)

	fmt.Printf("O tipo da variavel nome é: %T", nome)
	fmt.Println("")
}
