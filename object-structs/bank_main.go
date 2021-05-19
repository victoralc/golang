package main

import "fmt"

type Account struct {
	owner         string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
	// Creating object
	victorAccount := Account{owner: "Victor", agencyNumber: 589, accountNumber: 123456, balance: 4500.0}
	liviaAccount := Account{owner: "Livia", agencyNumber: 587, accountNumber: 3452421, balance: 3500.0}
	fmt.Println(victorAccount, liviaAccount)
}
