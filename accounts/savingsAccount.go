package accounts

import "github/banco/customers"

type SavingsAccount struct {
	Customer                   customers.Customer
	Agency, Account, Operation int
	balance                    float64
}

func (c *SavingsAccount) Withdraw(withdrawValue float64) (string, float64) {
	canBeWithdraw := withdrawValue <= c.balance && withdrawValue > 0

	if canBeWithdraw {
		c.balance -= withdrawValue
		return "Withdraw successfully", c.balance
	}
	return "No enough cash", c.balance
}

func (c *SavingsAccount) Deposit(deposit float64) (string, float64) {
	if deposit <= 0 {
		return "Value not accept", c.balance
	}
	c.balance += deposit
	return "Deposited successfully", c.balance
}

func (c *SavingsAccount) Transfer(transferValue float64, destinyAccount *SavingsAccount) bool {
	if c.balance < transferValue || transferValue < 0 {
		return false
	}
	c.balance -= transferValue
	destinyAccount.Deposit(transferValue)
	return true
}

func (c *SavingsAccount) GetBalance() float64 {
	return c.balance
}
