package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", count)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
	
}

func count(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("my-counter")

	if err == http.ErrNoCookie {
		log.Println("cookie not exist yet will be set now...")
		http.SetCookie(w, &http.Cookie{
			Name: "my-counter",
			Value: "0",
		})

	}

	// read cookie value and conv it to int
	v, _ := strconv.Atoi(c.Value)
	// after incrementing the value convert back to string
	v++
	c.Value = strconv.Itoa(v)

	http.SetCookie(w, c)

	fmt.Fprintln(w, "THE COOKIE IS SET - CHECK THE DEV TOOLS")
	fmt.Fprintf(w, "It is your %s visit", c.Value)
}
