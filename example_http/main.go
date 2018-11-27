package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handleHello)
	fmt.Println("serving on http://localhost:7777/hello")
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	log.Println("serving", req.URL)
	fmt.Fprintln(w, "Hello , world, this is the server as easy as possible. Ease of programming eeeease of programming")
}
