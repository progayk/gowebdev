package main

import (
	"github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
	"strings"
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
	v := c.Value
	vx := strings.Split(v, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", vx)
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("sID")
	if err != nil { // not nil if cookie doesn't exist
		// gen new ID
		uID, err := uuid.NewV4()
		if err != nil {
			log.Println("couldnt gen a an ID", err)
		}

		values := appendValues()
		v := uID.String() + "|" + values

		// create a cookie
		c = &http.Cookie{
			Name:  "sID",
			Value: v,
		}
		// set max age to 30 sec
		c.MaxAge = 10

		// write cookie to response
		http.SetCookie(w, c)
	}
	return c
}

func appendValues() (v string) {
	p1 := "black.jpg"
	p2 := "white.jpg"
	p3 := "orange.jpg"

	vx := []string{p1,p2,p3}
	v = strings.Join(vx, "|")
	return
}

