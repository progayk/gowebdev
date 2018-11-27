package main

import (
	"fmt"
	"sync"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	// send
	go send(eve, odd)

	// receive
	go receive(eve, odd, fanIn)

	for v := range fanIn {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func send(e, o chan<- int) {
	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

// receive acts as a receiver
func receive(e, o <-chan int, f chan<- int) {

	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		for v := range e {
			f <- v
		}
		wg.Done()
	}()

	go func(){
		for v := range o {
			f <- v
		}
		wg.Done()
	}()
	wg.Wait()

	close(f)
}

