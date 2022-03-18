package main

import "fmt"

func main() {
	firstName := "Victor"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Alcantara"
	fmt.Println(ptr, *ptr)
}