package main

import (
	"fmt"
	a "victor/golang/example_applications/banco/accounts"
	c "victor/golang/example_applications/banco/clients"
)

type verifyAccount interface {
	Withdraw(amount float64) string
}

func PayBill(account verifyAccount, amount float64) string{
	return account.Withdraw(amount)
}

func main() {
	victorAccount := a.Account{
		Holder:        c.AccountHolder{Name: "Victor Alcantara", Profession: "Software Engineer"},
		AgencyNumber:  1234,
		AccountNumber: 33456,
	}

	joaoAccount := a.Account{
		Holder:        c.AccountHolder{Name: "Victor Alcantara"},
		AgencyNumber:  1234,
		AccountNumber: 33456,
	}

	victorAccount.Deposit(1000)
	fmt.Println("Victor balance: ", victorAccount.GetBalance())
	fmt.Println("John balance: ", joaoAccount.GetBalance())

	fmt.Println("Transferring money...")
	victorAccount.Transfer(500, &joaoAccount)

	fmt.Println("Victor balance: ", victorAccount.GetBalance())
	fmt.Println("John balance: ", joaoAccount.GetBalance())

	fmt.Println("Victor Account - Pay bill")
	PayBill(&victorAccount, -400.0)

	fmt.Println("Victor balance after payment: ", victorAccount.GetBalance())
}
