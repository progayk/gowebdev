package main

import (
	"github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// check if any cookie exists in the request, if not create a new one
	c := getCookie(w, r)
	tpl.ExecuteTemplate(w, "index.gohtml", c.Value)
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("sID")
	if err != nil { // not nil if cookie doesn't exist
		// gen new ID
		uID, err := uuid.NewV4()
		if err != nil {
			log.Println("couldnt gen a an ID", err)
		}

		// create a cookie
		c = &http.Cookie{
			Name:  "sID",
			Value: uID.String(),
		}
		// set max age to 30 sec
		c.MaxAge = 10

		// write cookie to response
		http.SetCookie(w, c)
	}
	return c
}
