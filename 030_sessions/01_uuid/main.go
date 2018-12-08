package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	var id uuid.UUID
	var salt byte

	salt = 4
	id = uuid.NewMD5(id, []byte{salt})
	fmt.Println("id", id)

	c, err := r.Cookie("session")
	if err != nil {
		// means there is no cookie
		c = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true, // if HTTPS
			HttpOnly: true, // you can't access this cookie from javascript, it's http only
		}
		http.SetCookie(w, c)
	}
	fmt.Println(c)

}
