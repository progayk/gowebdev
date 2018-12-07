package main

import (
	"log"
	"net/http"
)

func main() {
	// When you have a index.html file on the root it will be served and not client.go.
	// However, when I make req on localhost:8080/client.go I can see the source code.
	// TODO: How is it excepted from the path?
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}


