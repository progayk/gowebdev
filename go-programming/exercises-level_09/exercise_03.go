package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Program has started")
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("goroutines", runtime.NumGoroutine())

	// create a var to be incremented
	counter := 0

	// number of goroutines
	const gs = 100

	// create a wait group
	var wg sync.WaitGroup

	// add goroutines
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			// read the value of counter
			v := counter
			// yield the processor
			runtime.Gosched()

			v++
			// write the value into to counter var
			counter = v

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
