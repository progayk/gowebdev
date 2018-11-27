package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Program has started")
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("goroutines", runtime.NumGoroutine())

	// create a var to be incremented
	var counter int64

	// number of goroutines
	const gs = 100

	// create a wait group
	var wg sync.WaitGroup

	// add goroutines
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			// read the value of counter
			atomic.AddInt64(&counter, 1)
			// yield the processor
			runtime.Gosched()
			// write the value into to counter var
			atomic.LoadInt64(&counter)
			// remove a goroutine from wait group
			wg.Done()
		}()
		fmt.Println("goroutines", runtime.NumGoroutine())
	}
	// wait till all goroutines to finish
	wg.Wait()

	fmt.Println("Counter", counter)

	fmt.Println("goroutines", runtime.NumGoroutine())
	fmt.Println("about the exit the program")
}

