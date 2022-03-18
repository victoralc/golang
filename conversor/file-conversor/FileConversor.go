package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Read JSON file
	data, err := os.ReadFile("transactions.json")
	if err != nil {
		panic(err)
	}

	// Get the content
	fmt.Print(string(data))

	var objectJson interface{}

	//Map string (json) to object
	errUnmarshall := json.Unmarshal(data, &objectJson)
	if errUnmarshall != nil {
		panic(errUnmarshall)
	}

	fmt.Print(objectJson)
	// Convert JSON to XML


}
