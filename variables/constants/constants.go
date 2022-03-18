package main

import (
	"fmt"
	"math"
)

const pi = 3.1415

const(
	number1 = iota + 10
	number2
	number3
	number4
	number5
)

const(
	number6 = iota
)

const s string = "string"

func main() {

	const i int = 14
	fmt.Println(i + 10)

	fmt.Println(float32(i) + 100.0)

	fmt.Println(number1, number2, number3, number4, number5, number6)

	// const number
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

}