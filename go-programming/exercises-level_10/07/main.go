package main

import "fmt"

func main() {
	c := make(chan int, 10)
	defer close(c)

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
		}()
	}

	for i := 0; i < 100; i++ {
		fmt.Println(i, <-c)
	}
}
