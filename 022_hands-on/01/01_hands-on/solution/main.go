package main

import (
	"fmt"
	"html/template"
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

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		d := data{
			Page: "index",
			Name: "",

		}
		err := tpl.ExecuteTemplate(w, "index.gohtml", d)
		HandleError(err)
	})

	mux.HandleFunc("/dog", func(w http.ResponseWriter, r *http.Request) {
		d := data{
			Page: "dog",
			Name: "",

		}
		err := tpl.ExecuteTemplate(w, "index.gohtml", d)
		HandleError(err)
	})

	mux.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		HandleError(err)

		d := data{
			Page: "me",
			Name: "Mayk",

		}

		err = tpl.ExecuteTemplate(w, "index.gohtml", d)
		HandleError(err)
	})


	err := http.ListenAndServe(":8080", mux)
	HandleError(err)
}
