package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	fmt.Println("func foo")
	wg.Done()
}

func bar() {
	fmt.Println("func bar")
	wg.Done()
}

func main() {
	wg.Add(2)


	go bar()
	go foo()
	//fmt.Println(runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("exiting the program...")
}
