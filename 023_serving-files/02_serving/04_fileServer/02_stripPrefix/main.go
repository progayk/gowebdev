package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/style/", http.StripPrefix("/style",http.FileServer(http.Dir("./style"))))
	http.Handle("/resources/", http.StripPrefix("/resources",http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<link href="style/style.css" rel="stylesheet" />
	<h1>Evet calisiyordu</h1>
	<img src="resources/toby.jpg" />`)
}

