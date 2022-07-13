package go_concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{}

	pool.Put("imam")
	pool.Put("ahmad")
	pool.Put("fahrezi")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Printf("data: %v\n", data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	fmt.Println("Selesai")
}
