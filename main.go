package main

import (
	"fmt"

	"github/banco/accounts"
	"github/banco/customers"
)

func PayBills(account verifyAccount, billsValue float64) {
	account.Withdraw(billsValue)
}

type verifyAccount interface {
	Withdraw(float64) (string, float64)
}

func main() {
	customerOne := customers.Customer{
		Name:       "Wanderson",
		Document:   "12345678910",
		Profession: "Desenvolvedor",
	}
	accountOne := accounts.SavingsAccount{
		Customer: customerOne,
		Agency:   0001,
		Account:  654321,
	}

	accountOne.Deposit(200)
	PayBills(&accountOne, 50)
	fmt.Println(accountOne.GetBalance())

}
