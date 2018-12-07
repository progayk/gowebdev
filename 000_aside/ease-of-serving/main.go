package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
		handleError(err)
	})
	router.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-Type", "text/html")
		err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
		handleError(err)
	})
	router.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
		handleError(err)
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleError(err error)  {
	if err != nil {
		log.Fatalln(err)
	}
}
