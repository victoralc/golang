package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words := os.Args[1:]

	report(frequencyOfFirstLetterIn(words))
}

func frequencyOfFirstLetterIn(words []string) map[string]int {
	frequency := make(map[string]int)
	for _, word := range words {
		first := strings.ToUpper(string(word[0]))
		counter, found := frequency[first]
		if found {
			frequency[first] = counter + 1
		} else {
			frequency[first] = 1
		}
	}
	return frequency
}

func report(frequency map[string]int) {
	fmt.Println("Frequency of first letter in words:")

	for first, counter := range frequency {
		fmt.Printf("%s = %d\n", first, counter)
	}
}
