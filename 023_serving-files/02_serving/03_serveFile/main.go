package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.HandleFunc("/style.css", style)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<link href="/style.css" rel="stylesheet" />
	<img src="/toby.jpg" />
	`)
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "toby.jpg")
}

func style(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "style.css")
}
