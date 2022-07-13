package go_concurrency

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("hello world")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("ups")
	time.Sleep(1 * time.Second)

}

func DisplayNumber(num int) {
	fmt.Println("number: ", num)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

}
