package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "my-cookie",
		Value: "some value from mayk",
	})
	fmt.Fprint(w, "COOKIE WRITTEN - CHECK YOUR BROWER'S DEV TOOLS")
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	fmt.Println("the cookie domain is ", c.Domain)
	fmt.Println("the cookie path is ", c.Path)
	fmt.Println("the cookie max age is ", c.MaxAge)
	fmt.Println("the cookie expires is ", c.Expires)
	fmt.Println("the cookie secure is ", c.Secure)
	fmt.Fprint(w, "the cookie is ", c)
}
