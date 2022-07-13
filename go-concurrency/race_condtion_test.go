package go_concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	counter := 0
	mutex := sync.Mutex{}

	for i := 0; i < 1000; i++ {

		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				counter = counter + 1
			}

		}()

	}

	fmt.Println(counter)
	time.Sleep(5 * time.Second)
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()

}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()

}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount

}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "imam",
		Balance: 1000000,
	}
	user2 := UserBalance{
		Name:    "ihab",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 100000)
	time.Sleep(3 * time.Second)

	fmt.Println("User", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User", user2.Name, ", Balance ", user2.Balance)
}

func TestAtomic(t *testing.T) {
	var counter int32 = 0

	for i := 0; i < 1000; i++ {

		go func() {
			for j := 0; j < 100; j++ {
				atomic.AddInt32(&counter, 1)
			}

		}()

	}

	time.Sleep(5 * time.Second)
	fmt.Println(counter)
}
