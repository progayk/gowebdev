package main

import (
	"github.com/google/uuid"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type user struct {
	UserName string
	First string
	Last string
}

var dbUsers = map[string]user{} // user ID, user
var dbSessions = map[string]string{} // session= ID, user ID

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	sID, _ := uuid.NewUUID()

	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		// means there is no cookie
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
			// Secure: true, // if HTTPS
			HttpOnly: true, // you can't access this cookie from javascript, it's http only
		}
		http.SetCookie(w, c)
	}

	var u user

	// if the user exists already, get the user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	// process form submission
	if r.Method == http.MethodPost {
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		un := r.FormValue("username")

		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)

}

func bar(w http.ResponseWriter, r *http.Request) {
	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// get user if exists, else redirect
	u, ok := dbUsers[dbSessions[c.Value]]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
