package go_concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Printf("time.Now(): %v\n", time.Now())

	time := <-timer.C
	fmt.Printf("time: %v\n", time)
}

func TestTimerAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Printf("time.Now(): %v\n", time.Now())

	time := <-channel
	fmt.Printf("time: %v\n", time)
}

func TestAfterFunc(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		wg.Done()
	})
	fmt.Println(time.Now())
	wg.Wait()
}
