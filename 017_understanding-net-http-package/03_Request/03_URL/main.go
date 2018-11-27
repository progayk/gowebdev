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
		Submissions map[string][]string
		URL *url.URL
		Header http.Header
		Host string
		ContentLength int64
	}{
		Method: r.Method,
		Submissions: r.Form,
		URL: r.URL,
		Header: r.Header,
		Host: r.Host,
		ContentLength: r.ContentLength,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func main() {
	var handler hotdog

	http.ListenAndServe(":8080", handler)
}