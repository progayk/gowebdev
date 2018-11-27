package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go routines\t", runtime.NumGoroutine())
	c := make(chan int)

	// send
	go func(){
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	fmt.Println("Go routines\t", runtime.NumGoroutine())

	// receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about the exit")
}
