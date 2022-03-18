package main

import "fmt"

func main() {

	//using for with external i
	i := 1
	for i <= 10 {
		fmt.Println("Printing: ", i)
		i++
	}

	// using declaring variables within
	for j := 0; j < 10; j++ {
		fmt.Println("Printing: ", j)
	}

	// with break
	for {
		fmt.Println("loop")
		break
	}

	// using continue
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	numbersDividedBy5BetweenZeroAndFifty()

}

func numbersDividedBy5BetweenZeroAndFifty()  {
	for i := 0; i <= 50; i++ {
		if i%5 == 0{
			fmt.Println(i, " is a number divided by 5")
		}
	}
}
