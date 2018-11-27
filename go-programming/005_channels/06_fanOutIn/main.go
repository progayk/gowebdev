package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)
	go fanInOut(c1, c2)

	xi := []int{}
	for v := range c2 {
		fmt.Println(v)
		xi = append(xi, v)
	}

	fmt.Println("Unsorted array of numbers", xi)

}

func populate(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanInOut(c1, c2 chan int) {

	var wg sync.WaitGroup
	//wg.Add(100)
	for v := range c1 {
		wg.Add(1)
		go func(v int) {
			t := timeConsumingTask(v)
			c2 <- t
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)

}

func timeConsumingTask(n int) int {
	time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
	return n
}
