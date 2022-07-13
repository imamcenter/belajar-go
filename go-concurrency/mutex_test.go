package go_concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	counter := 0
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				counter = counter + 1

				mutex.Unlock()
			}
		}()

	}
	time.Sleep(5 * time.Second)
	fmt.Println(counter)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (a *BankAccount) AddBalance(amount int) {
	a.RWMutex.Lock()
	a.Balance = a.Balance + amount
	a.RWMutex.Unlock()
}

func (a *BankAccount) GetBalance() int {
	a.RWMutex.RLock()
	balance := a.Balance
	a.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	bankAccount := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				bankAccount.AddBalance(1)
				fmt.Println(bankAccount.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)

}
