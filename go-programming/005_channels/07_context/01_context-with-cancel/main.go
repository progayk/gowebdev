package _1_context_with_cancel

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, close := context.WithCancel(context.Background())
	defer close() // defer the close to run before the program exit

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // returning not to leak goroutine
			case dst <- n:
				time.Sleep(1000 * time.Millisecond)
				n++
			}
		}
	}()
	return dst
}

