package main

import (
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>doggy doggy doggy</h1>")
}

type hotcat int

func (h hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>kitty kitty kitty</h1>")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
