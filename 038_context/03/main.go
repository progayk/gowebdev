package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, "userID", 007)
	ctx = context.WithValue(ctx, "fname", "James")

	res, err := accessDB(ctx)
	if err != nil {
		log.Panic("LOOK AT ME ", err)
	}

	fmt.Fprintln(w, res)
}

func accessDB(ctx context.Context) (int, error) {

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	ch := make(chan int)

	go func() {
		// ridiculously long running task
		uID := ctx.Value("userID").(int)
		time.Sleep(10*time.Second)

		// check to make sure we are not running in vain
		// if ctx.Done() has
		if ctx.Err() != nil {
			return
		}

		ch <- uID
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Println(ctx)
	fmt.Fprintln(w, ctx)
}
