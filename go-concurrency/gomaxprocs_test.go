package go_concurrency

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGomaxprocs(t *testing.T) {
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())

	runtime.GOMAXPROCS(1)
	fmt.Printf("runtime.GOMAXPROCS(0): %v\n", runtime.GOMAXPROCS(0))

	fmt.Printf("runtime.NumGoroutine(): %v\n", runtime.NumGoroutine())
}
