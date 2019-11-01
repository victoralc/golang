package main

import (
	"fmt"
	"os"
	"strconv"
)

func printPersonData(name string, age int) {
	fmt.Printf("%s is %d years old.\n", name, age)
}

func sum(n, m int) int {
	return n + m
}

func printSum(n, m int) {
	fmt.Printf("Sum: %d + %d = %d", n, m, sum(n, m))
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Missing arguments: <name> <age>")
		os.Exit(1)
	}

	name := os.Args[1]
	age := os.Args[len(os.Args)-1]

	ageInt, err := strconv.Atoi(age)
	if err != nil {
		fmt.Printf("Age: %d is invalid!\n", ageInt)
		os.Exit(1)
	}

	printPersonData(name, ageInt)
	printSum(40, 20)
}
