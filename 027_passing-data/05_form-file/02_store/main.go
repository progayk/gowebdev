package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	var s string
	if r.Method == http.MethodPost {
		// get form data
		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		defer f.Close()

		// for your information
		fmt.Println("\nfile", f, "\nfile header", h, "\nerror", err)

		// read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		s = string(bs)

		// store on server
		dts, err := os.Create(filepath.Join("./user/", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dts.Close()

		_, err = dts.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	w.Header().Set("Content-Type", "text/html; charset: utf-8")
	tpl.ExecuteTemplate(w, "index.gohtml", s)

}
