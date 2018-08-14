package pointers

import (
	"errors"
	"fmt"
)

//BitCoin as currency
type BitCoin int

//Wallet bitcoin
type Wallet struct {
	balance BitCoin
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//Deposit the bitcoin
func (w *Wallet) Deposit(amount BitCoin) {
	w.balance += amount
}

//Balance check the balance
func (w *Wallet) Balance() BitCoin {
	return w.balance
}

//Widraw to wdiraw bitcoings
func (w *Wallet) Widraw(amt BitCoin) error {
	if w.balance < amt {
		return errors.New("insufficient balance")
	}
	w.balance -= amt
	return nil
}
