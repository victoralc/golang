package accounts

import "victor/golang/object-structs/clients"

type SavingsAccount struct {
	Client clients.Client
	AccountNumber, AgencyNumber, Operation int
	balance       float64
}

func (a *SavingsAccount) Withdraw(value float64) string {
	canWithdraw := value > 0 && value <= a.balance
	if canWithdraw {
		a.balance -= value
		return "Withdraw completed."
	} else {
		return "Insufficient balance."
	}
}

func (a *SavingsAccount) Deposit(value float64) (string, float64){
	if value > 0 {
		a.balance += value
		return "Deposit completed. Your new balance:", a.balance
	} else {
		return "Invalid value. Your balance:", a.balance
	}
}

func (a *SavingsAccount ) Transfer(valueToBeTransferred float64, destinyAccount *Account) bool {
	if valueToBeTransferred < a.balance && valueToBeTransferred > 0 {
		a.balance -= valueToBeTransferred
		destinyAccount.Deposit(valueToBeTransferred)
		return true
	}
	return false
}

func (a *SavingsAccount) GetBalance() float64 {
	return a.balance
}


