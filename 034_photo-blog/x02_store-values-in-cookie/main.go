package main

import (
	"fmt"
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
	v := values(c)
	tpl.ExecuteTemplate(w, "index.gohtml", v)
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("sID")
	if err != nil { // not nil if cookie doesn't exist
		// gen new ID
		uID, err := uuid.NewV4()
		if err != nil {
			log.Println("couldnt gen a an ID", err)
		}

		values := "|black|white|orange"
		v := uID.String() + values

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

func values(c *http.Cookie) []string {
	xv := strings.FieldsFunc(c.Value, func(c rune) bool {
		s := []rune("|")
		return c == s[0]
	})
	fmt.Println("slice of values", xv)

	return xv
}