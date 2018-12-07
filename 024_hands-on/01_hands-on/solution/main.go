package main

import (
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func foo(w http.ResponseWriter, r *http.Request)  {
	r.Header.Set("Content-Type", "text/plain")
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	if err := tpl.ExecuteTemplate(w, "dog.gohtml", nil); err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
	}
}

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(foo))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.ListenAndServe(":8080", nil)
}