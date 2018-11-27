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
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}
