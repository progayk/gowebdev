package main

import (
	"fmt"
	"runtime"
	"sync"
)

// run this program with race flag => go run -race mj-testing.go

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// Gosched func makes goroutine to run seperately, don't wait for it to be done
			runtime.Gosched()
			v++
			counter = v

			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("------------")
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
