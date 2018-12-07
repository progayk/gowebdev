package main

import "fmt"

func main() {

	// this won't run since we have 1 room for a value.
	c := make(chan int, 1)

	c <- 42
	c <- 43

	fmt.Println(<-c)
}

