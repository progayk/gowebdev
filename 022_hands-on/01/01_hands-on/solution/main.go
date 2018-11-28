package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type data struct{
	Page string
	Name string
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		d := data{
			Page: "index",
			Name: "",

		}
		err := tpl.ExecuteTemplate(w, "index.gohtml", d)
		if err != nil {
			log.Println(err)
		}
	})

	mux.HandleFunc("/dog", func(w http.ResponseWriter, r *http.Request) {
		d := data{
			Page: "dog",
			Name: "",

		}
		err := tpl.ExecuteTemplate(w, "index.gohtml", d)
		if err != nil {
			log.Println(err)
		}
	})

	mux.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		d := data{
			Page: "me",
			Name: "Mayk",

		}

		err = tpl.ExecuteTemplate(w, "index.gohtml", d)
		if err != nil {
			log.Println(err)
		}
	})


	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
	}


}
