package main

import (
	"github.com/google/uuid"
	"html/template"
	"net/http"
)

var tpl *template.Template

type user struct {
	UserName string
	Password string
	First string
	Last string
}

var dbUsers = map[string]user{} // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	tpl.ExecuteTemplate(w, "index.gohtml", u)

}

func bar(w http.ResponseWriter, r *http.Request) {
	if !isLoggedIn(r) {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}

	c, _ := r.Cookie("session")
	// find user
	u := dbUsers[dbSessions[c.Value]]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)

}

func signup(w http.ResponseWriter, r *http.Request) {
	if isLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		// process form submission
		un := r.FormValue("username")
		pw := r.FormValue("password")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")

		// create a new user
		u := user{un, pw, f, l}


		// create a new cookie
		sID, _ := uuid.NewUUID()
		c := &http.Cookie{
			Name: "session",
			Value: sID.String(),
			HttpOnly: true,
		}

		// store values on db
		dbSessions[c.Value] = un
		dbUsers[un] = u

		// redirect
		http.SetCookie(w, c)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)

}
