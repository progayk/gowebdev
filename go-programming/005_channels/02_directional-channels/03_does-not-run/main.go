package main

import "fmt"

func main() {

	// Does not run since it's a receive-only type chan
	c := make(<-chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("---------")
	fmt.Printf("%T\n", c)
}

