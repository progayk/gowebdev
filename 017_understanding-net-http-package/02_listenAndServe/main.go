package main

import (
	"fmt"
	"net/http"
)

type hotdog string

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// since hotdog has a method ServeHTTP attached to it,
	// it's of type http.Handler, as well
	rm := req.Method
	rh := req.Header
	fmt.Printf("%T\t%v\n", *req, rm)
	for k, v := range rh {
		fmt.Println(k, v)
	}
	w.Header().Set("mj-header", "mayk" )
	fmt.Fprintln(w, "<h1>any function you want to pass in here</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
