package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	var s string
	if r.Method == http.MethodPost {
		f, fh, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
		}
		defer f.Close()

		fmt.Println("\nfile", f, "\nheader", fh, "\nerror", err)

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}

	w.Header().Set("Content-Type", "text/html; charset: utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
		<input type="file" name="q">
		<input type="submit">
	</form><br>
	`+s)


}
