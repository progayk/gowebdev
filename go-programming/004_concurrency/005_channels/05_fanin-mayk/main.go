package main

import (
	"fmt"
	"sort"
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

	var xi []int
	for v := range fanIn {
		xi = append(xi, v)
	}

	// create an array with the coming numbers
	// sort this array
	fmt.Println("Unsorted: written as it comes")
	fmt.Println(xi)

	fmt.Println("Sorted")
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("about to exit")
}

func send(e, o chan<- int) {
	// send off two channels
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		for i := 0; i < 100; i++ {
			if i % 2 == 0 {
				e <- i
			}
		}
		wg.Done()
	}()

	go func(){
		for i := 0; i < 100; i++ {
			if i % 2 != 0 {
				o <- i
			}
		}
		wg.Done()
	}()
	wg.Wait()

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

