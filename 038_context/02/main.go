package main

import (
	"context"
	"fmt"
	"net/http"
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

	res := accessDB(ctx)

	fmt.Fprintln(w, res)
}

func accessDB(ctx context.Context) int {
	uID := ctx.Value("userID").(int)
	return uID
}

func bar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Println(ctx)
	fmt.Fprintln(w, ctx)
}
