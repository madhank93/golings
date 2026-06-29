// pointers2
// Passing a *struct lets a function mutate the caller's struct. Go
// auto-dereferences, so you write a.Field even when a is a pointer.

// I AM NOT DONE
package main_test

import "testing"

type Account struct {
	Balance int
}

// deposit adds amount to the account's balance, mutating the original.
func deposit(a *Account, amount int) {
	// FIXME: add amount to the balance via the pointer.
}

func TestDeposit(t *testing.T) {
	acc := Account{Balance: 100}
	deposit(&acc, 50)
	if acc.Balance != 150 {
		t.Errorf("want 150, got %d", acc.Balance)
	}
}
