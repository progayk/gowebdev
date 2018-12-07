package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type data struct {
	Page string
	Name string
}

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	d := data{"index", ""}
	err := tpl.ExecuteTemplate(w, "index.html", d)
	HandleError(err)
}

func dog(w http.ResponseWriter, r *http.Request) {
	d := data{"dog", ""}
	err := tpl.ExecuteTemplate(w, "index.html", d)
	HandleError(err)
}

func me(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	HandleError(err)
	fmt.Println(r.Form)

	f := r.Form
	name := strings.Join(f["name"], "")


	fmt.Println(name)
	fmt.Printf("type of f['name'] = %T\n", name)

	d := data{"me", name}

	err = tpl.ExecuteTemplate(w, "index.html", d)
	HandleError(err)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(index))
	mux.Handle("/dog", http.HandlerFunc(dog))
	mux.Handle("/me", http.HandlerFunc(me))

	// visit: http://localhost:8080/me?name=Mayk
	err := http.ListenAndServe(":8080", mux)
	HandleError(err)
}
