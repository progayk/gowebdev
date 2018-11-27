package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send
	go send(eve, odd, quit)

	// receive
	receive(eve, odd, quit)

	fmt.Println("about to exit")
}

// receive acts as a receiver
func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("the value received by the eve channel: ", v)
		case v := <-o:
			fmt.Println("the value received by the odd channel: ", v)
		case v := <-q:
			fmt.Println("the value received by the quit channel: ", v)
			return
		}
	}

}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- true
}
