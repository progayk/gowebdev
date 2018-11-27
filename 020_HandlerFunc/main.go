package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>doggy doggy doggy</h1>")
}

func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>kitty kitty kitty</h1>")
}

func main() {
	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)
}

// this is similar to this:
// https://play.golang.org/p/X2dlgVSIrd
// ---and this---
// https://play.golang.org/p/YaUYR63b7L`

