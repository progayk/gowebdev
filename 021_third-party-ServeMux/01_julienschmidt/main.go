package main

import (
	"fmt"
	"html/template"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/contact", contact)
	router.GET("/about", about)
	router.GET("/apply", apply)
	router.POST("/apply", applyProcess)
	router.GET("/user/:name", user)
	router.GET("/blog/:category/:article", blogRead)
	router.POST("/blog/:category/:article", blogWrite)
	http.ListenAndServe(":8080", router)
	
}

func blogRead(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Blog READ Category: %s\n", ps.ByName("category"))
	fmt.Fprintf(w, "Blog READ Article: %s\n", ps.ByName("article"))
	fmt.Printf(r.RequestURI)
}

func blogWrite(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Blog WRITE Category: %s\n", ps.ByName("category"))
	fmt.Fprintf(w, "Blog WRITE Article: %s\n", ps.ByName("article"))
}

func user(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "USER is %s\n", ps.ByName("name"))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func applyProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
