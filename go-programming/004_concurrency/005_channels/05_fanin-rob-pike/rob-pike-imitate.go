package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("yosef"), boring("kirk"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("You both are boring, I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			m := fmt.Sprintf("%d. %s", i, msg)
			c <- m
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			// input1: it receives from a channel and put it into c channel
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- input2
		}
	}()
	return c
}
