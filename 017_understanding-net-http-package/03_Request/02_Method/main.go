package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct{
		Method string
		Submissions url.Values
	}{
		Method: r.Method,
		Submissions: r.Form,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func main() {
	var handler hotdog

	http.ListenAndServe(":8080", handler)
}