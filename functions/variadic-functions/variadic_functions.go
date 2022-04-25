package main

import (
	"fmt"
	"os"
)

func CreateFiles(baseDir string, files ...string) {
	for _, name := range files {
		pathFile := fmt.Sprintf("%s/%s.%s", baseDir, name, "txt")

		file , err := os.Create(pathFile)

		defer file.Close()

		if err != nil {
			fmt.Printf("Error occurred during file creation %s: %v\n", name, err)
			os.Exit(1)
		}

		fmt.Printf("File %s created. \n", file.Name())
	}
}

func main()  {
	tmp := os.TempDir()

	CreateFiles(tmp)
	CreateFiles(tmp, "test1")
	CreateFiles(tmp, "test2", "test3", "test4")
}
