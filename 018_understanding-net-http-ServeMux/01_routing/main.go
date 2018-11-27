package main

import (
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		io.WriteString(w, "<h1>doggy doggy doggy</h1>")
	case "/cat":
		io.WriteString(w, "catty catty")
	default:
		io.WriteString(w, "mayki mayki mayki")
	}
}

func main() {
	var handler hotdog

	http.ListenAndServe(":8080", handler)
}
