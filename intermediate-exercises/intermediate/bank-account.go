package intermediate

import (
	"errors"
	"fmt"
	"sync"
)

type BankAccount struct {
	mu      sync.Mutex
	owner   string
	balance float64
}

func (b *BankAccount) Deposit(amount float64) error {
	switch {
	case amount < 0:
		return errors.New("Amount must be positive")
	default:
		b.mu.Lock()
		defer b.mu.Unlock()
		b.balance += amount
		return nil
	}
}

func (b *BankAccount) Withdraw(amount float64) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	switch {
	case amount > b.balance:
		return errors.New("Not enough balance")
	default:
		b.balance -= amount
		return nil
	}
}

func (b *BankAccount) Balance() float64 {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.balance
}

func (b *BankAccount) Owner() string {
	return b.owner
}

func NewBankAccount(owner string) *BankAccount {
	return &BankAccount{owner: owner}
}

func ExecuteBankAccount() {
	account := NewBankAccount("Alice")

	err := account.Deposit(100)
	if err != nil {
		fmt.Println("Deposit error:", err)
	}

	err = account.Withdraw(30)
	if err != nil {
		fmt.Println("Withdraw error:", err)
	}

	fmt.Printf("%s's balance: %.2f\n", account.Owner(), account.Balance())

	err = account.Withdraw(1000) // Should return an error
	if err != nil {
		fmt.Println("Withdraw error:", err)
	}
}
