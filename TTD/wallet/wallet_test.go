package main

import(
	"testing"
)

func TestWallet(t *testing.T){
	assertBalence := func(t testing.TB, w Wallet, hasBalence coin){
		t.Helper()	
		got := w.getBalence()
		if got != hasBalence {
			t.Errorf("got %s has balence %s", got, hasBalence)
		}
	}
	assertError := func(t testing.TB, got, want error){
		t.Helper()
		if got == nil{
			t.Fatal("wanted an error but didn't get one")
		}
		if got != want{
			t.Errorf("got %q, want %q", got, want)
		}
	}
	assertNoError := func(t testing.TB, got error){
		t.Helper()
		if got != nil{
			t.Fatal("got and error but didn't want one")
		}
	}
	t.Run("deposit",func(t *testing.T){
		w := Wallet{}
		w.Deposit(coin(10))
		assertBalence(t, w, coin(10)) 	
	})
	t.Run("withdraw",func(t *testing.T){
		w := Wallet{}
		w.Deposit(coin(20))
		err:=w.Withdraw(coin(5))
		assertNoError(t, err)
		assertBalence(t, w, coin(15))

	})
	t.Run("unsufficient funds", func(t *testing.T){
		startBal := coin(20) 
		w := Wallet{Balence: startBal}	
		err := w.Withdraw(coin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalence(t, w, coin(startBal))
	})
}//testWallet

