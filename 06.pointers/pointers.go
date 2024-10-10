package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}
type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(amount Bitcoin) error {
	if amount <= 0 {
		return errors.New("the amount must be greater than zero")
	}
	w.balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount <= 0 {
		return errors.New("the amount must be greater than zero")
	}
	if amount > w.balance {
		return errors.New("withdrawal amount can't exceed the current balance")
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
