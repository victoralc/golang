package main

import "fmt"

func main() {

	number1 := 20
	number2 := 20

	if number1 > number2 {
		fmt.Printf("Number1: %d is greater than Number2: %d", number1, number2)
	} else if number1 == number2 {
		fmt.Printf("Number1: %d is equal than Number2: %d", number1, number2)
	} else {
		fmt.Printf("Number1: %d is shorter than Number2: %d", number1, number2)
	}

}