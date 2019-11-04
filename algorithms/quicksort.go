package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1:]
	numbers := make([]int, len(input))

	for i, n := range input {
		number, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s is not a valid number!\n", n)
			os.Exit(1)
		}
		numbers[i] = number

	}

	fmt.Println(quicksort(numbers))
}

func quicksort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	n := make([]int, len(numbers))
	copy(n, numbers)

	indexPivot := len(n) / 2
	pivot := n[indexPivot]

	n = append(n[:indexPivot], n[indexPivot+1:]...)

	shortens, highers := partition(n, pivot)

	return append(append(quicksort(shortens), pivot), quicksort(highers)...)
}

func partition(numbers []int, pivot int) (shortens []int, highers []int) {
	for _, n := range numbers {
		if n <= pivot {
			shortens = append(shortens, n)
		} else {
			highers = append(highers, n)
		}
	}
	return shortens, highers
}
