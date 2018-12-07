package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
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
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	// Remember: Always close watch you have opened.
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	http.ServeContent(w, r, f.Name(), fi.ModTime(), f)
	// TAKE-AWAY: vim mode R -> continuous re-write
}

func style(w http.ResponseWriter, r *http.Request) {
	css, err := os.Open("style.css")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	// Remember: Always close watch you have opened.
	defer css.Close()

	fi, err := css.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	fmt.Println(fi.ModTime())

	http.ServeContent(w, r, css.Name(), fi.ModTime(), css)
	// TAKE-AWAY: vim mode R -> continuous re-write
}
