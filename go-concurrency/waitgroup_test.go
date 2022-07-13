package go_concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsyncronous(wg *sync.WaitGroup) {
	m := sync.Mutex{}
	m.Lock()
	defer wg.Done()
	m.Unlock()
	wg.Add(1)
	fmt.Println("Hello")
	time.Sleep(2 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {

		go RunAsyncronous(group)
	}

	group.Wait()
	fmt.Printf("\"Selesai\": %v\n", "Selesai")
}
