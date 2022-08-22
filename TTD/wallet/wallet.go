package main

import(
	"fmt"
	"errors"
)

type coin int

type Wallet struct{
	Balence coin 
}//Wallet

type Stringer interface{
	String() string
}//Stringer

func (c coin) String() string{
	return fmt.Sprintf("%d BTC",c)
}

func (w *Wallet) Deposit(amount coin){
	w.Balence += amount 
}//Deposit

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficent funds")

func (w *Wallet) Withdraw(amount coin) error{
	if amount > w.Balence{
		return ErrInsufficientFunds
	}
	w.Balence -= amount
	return nil
}//Withdraw

func (w *Wallet) getBalence() coin{
	return w.Balence
}//getBalence
