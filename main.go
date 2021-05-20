package main

import (
	"fmt"
	"victor/golang/object-structs/accounts"
	"victor/golang/object-structs/clients"
)

func PayBill(account verifyAccount, value float64) {
	account.Withdraw(value)
}

type verifyAccount interface {
	Withdraw(value float64) string
}

func main() {

	// Creating object

	// Different memory addresses
	johnAccount := accounts.Account{
		Client: clients.Client{Name: "John", ID: "123456", Profession: "Software Engineer"},
		AgencyNumber:881,
		AccountNumber: 123456}

	johnAccount.Deposit(1000)

	fmt.Println("Account Owner:", johnAccount.Client.Name)
	fmt.Println("Initial balance:", johnAccount.GetBalance())
	fmt.Println("Withdraw cash of 300 ...")
	fmt.Println(johnAccount.Withdraw(300))
	message, balance := johnAccount.Deposit(10000)
	fmt.Println(message, balance)
	fmt.Println("Final balance:", johnAccount.GetBalance())
	fmt.Println(johnAccount)

	adrianAccount := accounts.SavingsAccount{Client: clients.Client{Name: "Adrian"}}
	adrianAccount.Deposit(3500)

	PayBill(&adrianAccount, 400)

	fmt.Println(adrianAccount.GetBalance())
}
