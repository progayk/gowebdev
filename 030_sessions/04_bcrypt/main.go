package main

import (
	"github.com/google/uuid"
	"html/template"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var tpl *template.Template

type user struct {
	UserName string
	Password []byte
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

		// username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "username already taken", http.StatusForbidden)
			return
		}

		// create session
		sID, _ := uuid.NewUUID()
		c := &http.Cookie{
			Name: "session",
			Value: sID.String(),
			HttpOnly: true,
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un

		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		u := user{un, bs, f, l}
		dbUsers[un] = u

		// redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}
