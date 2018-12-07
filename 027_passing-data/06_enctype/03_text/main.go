package main

import (
	"html/template"
	"log"
	"net/http"
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

	bs := make([]byte, r.ContentLength)
	r.Body.Read(bs)
	body := string(bs)

	//if r.Method == http.MethodPost {
	//	f, h, err := r.FormFile("q")
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		log.Fatalln(err)
	//	}
	//	defer f.Close()
	//
	//
	//}

	err := tpl.ExecuteTemplate(w, "index.gohtml", body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}

}
