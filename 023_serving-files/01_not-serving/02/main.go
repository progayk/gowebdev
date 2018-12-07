package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {

	r.Header.Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<!-- we don't serve from the server-->
	<img src="/toby.jpg" width=300/>
	`)
}
