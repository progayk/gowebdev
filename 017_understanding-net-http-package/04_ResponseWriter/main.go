package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("header-mayk", "mayks special header")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Any function you want to pass here</h1>")
}

func main() {
	var handler hotdog

	http.ListenAndServe(":8080", handler)
}
