package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, `<a href="/set">set a cookie</a>`)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "yusufs-cookie",
		Value: "yusufa ozel cookie degeri",
		Path: "/",
	})
	http.SetCookie(w, &http.Cookie{
		Name: "counter",
		Value: "0",
		Path: "/",
	})
	http.SetCookie(w, &http.Cookie{
		Name: "a",
		Value: "b",
		Path: "/",
	})
	fmt.Fprint(w, `<h1><a href="/read">Read the cookie</a></h1>`)
}

func read(w http.ResponseWriter, r *http.Request)  {
	c1, err := r.Cookie("yusufs-cookie")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}

	c2, err := r.Cookie("counter")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}

	count, err := strconv.Atoi(c2.Value)
	if err != nil {
		log.Print(err)
	}
	count++
	c2.Value = strconv.Itoa(count)

	http.SetCookie(w, c2)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<h2>It is your %v. visit</h2> yusufs cookie: %v<h1><a href="/expire">expire the cookie</a></h1>`, c2.Value, c1.Value)
}

func expire(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("yusufs-cookie")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}

	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

