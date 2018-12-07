package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	handleError(err)
}

func about(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	handleError(err)
}

func contact(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	handleError(err)
}

func apply(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
		handleError(err)
	case http.MethodPost:
		err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
		handleError(err)
	default:
		io.WriteString(w, "only GET and POST methods accepted.")
	}
}
