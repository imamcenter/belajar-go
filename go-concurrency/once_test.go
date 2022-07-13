package go_concurrency

import (
	"fmt"
	"sync"
	"testing"
)

var counter int = 0

func OnlyOne() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			once.Do(OnlyOne)
			// OnlyOne()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("counter: %v\n", counter)
}
