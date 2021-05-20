package accounts

import "victor/golang/object-structs/clients"

type Account struct {
	Client        clients.Client
	AgencyNumber  int
	AccountNumber int
	balance       float64
}

func (a *Account) Withdraw(value float64) string {
	canWithdraw := value > 0 && value <= a.balance
	if canWithdraw {
		a.balance -= value
		return "Withdraw completed."
	} else {
		return "Insufficient balance."
	}
}

func (a *Account) Deposit(value float64) (string, float64){
	if value > 0 {
		a.balance += value
		return "Deposit completed. Your new balance:", a.balance
	} else {
		return "Invalid value. Your balance:", a.balance
	}
}

func (a *Account ) Transfer(valueToBeTransferred float64, destinyAccount *Account) bool {
	if valueToBeTransferred < a.balance && valueToBeTransferred > 0 {
		a.balance -= valueToBeTransferred
		destinyAccount.Deposit(valueToBeTransferred)
		return true
	}
	return false
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

