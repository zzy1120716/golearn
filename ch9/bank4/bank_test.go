package bank_test

import (
	"fmt"
	"golearn/ch9/bank4"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// First withdraw
	go func() {
		if ok := bank.Withdraw(400); !ok {
			fmt.Println("insufficient balance...")
		}
		done <- struct{}{}
	}()

	// Second withdraw
	go func() {
		if ok := bank.Withdraw(200); !ok {
			fmt.Println("insufficient balance...")
		}
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done
	<-done
	<-done

	if got, want := bank.Balance(), 100; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
