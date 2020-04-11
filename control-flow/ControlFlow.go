package main

import "fmt"

func main() {
	fmt.Println("Hello, control flow")
	foo()
	fmt.Println("Another print here")

	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func bar() {
	fmt.Println("Bar() method calling ...")
}

func foo() {
	fmt.Println("Function Foo() calling..")
}
