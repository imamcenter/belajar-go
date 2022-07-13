package gocontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Printf("background: %v\n", background)

	todo := context.TODO()
	fmt.Printf("todo: %v\n", todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")
	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")
	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Printf("contextA: %v\n", contextA)
	fmt.Printf("contextB: %v\n", contextB)
	fmt.Printf("contextC: %v\n", contextC)
	fmt.Printf("contextD: %v\n", contextD)
	fmt.Printf("contextE: %v\n", contextE)
	fmt.Printf("contextF: %v\n", contextF)
	fmt.Printf("contextG: %v\n", contextG)

	fmt.Printf("contextF.Value(\"f\"): %v\n", contextF.Value("f"))
	fmt.Printf("contextF.Value(\"c\"): %v\n", contextF.Value("c"))
	fmt.Printf("contextF.Value(\"b\"): %v\n", contextF.Value("b"))
	fmt.Printf("contextA.Value(\"b\"): %v\n", contextA.Value("b"))
}

func CreateCounter() chan int {
	dst := make(chan int)

	go func() {
		defer close(dst)
		counter := 0
		for {
			dst <- counter
			counter++
		}
	}()

	return dst
}

func TestGoroutineLeak(t *testing.T) {
	fmt.Println("total goroutine", runtime.NumGoroutine())

	dst := CreateCounter()
	for n := range dst {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}
	time.Sleep(1 * time.Second)
	fmt.Println("total goroutine", runtime.NumGoroutine())

}

func CreateCounter1(ctx context.Context) chan int {
	dst := make(chan int)

	go func() {
		defer close(dst)
		counter := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				dst <- counter
				counter++
			}
		}
	}()

	return dst
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("total goroutine", runtime.NumGoroutine())
	parent := context.Background()

	ctx, cancel := context.WithCancel(parent)

	dst := CreateCounter1(ctx)
	for n := range dst {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}
	cancel()
	time.Sleep(5 * time.Second)
	fmt.Println("total goroutine", runtime.NumGoroutine())

}

func CreateCounter2(ctx context.Context) chan int {
	dst := make(chan int)

	go func() {
		defer close(dst)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				dst <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return dst
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("total goroutine", runtime.NumGoroutine())
	parent := context.Background()

	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	dst := CreateCounter2(ctx)
	for n := range dst {
		fmt.Println("counter", n)
	}
	fmt.Println("total goroutine", runtime.NumGoroutine())

}
func TestContextWithDeadline(t *testing.T) {
	fmt.Println("total goroutine", runtime.NumGoroutine())
	parent := context.Background()

	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	dst := CreateCounter2(ctx)
	for n := range dst {
		fmt.Println("counter", n)
	}
	fmt.Println("total goroutine", runtime.NumGoroutine())

}
