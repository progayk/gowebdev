package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request)  {
	v := r.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// change the method to get to see the form value in url
	io.WriteString(w, `
	<form method="post" action="/">
		<input type="text" name="q" />
		<input type="submit" value="submit"/>
	</form>
	`+v)

}