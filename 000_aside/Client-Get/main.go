package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	c := http.Client{}
	r, err := c.Get("https://www.google.com")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("status code:", r.StatusCode)
	for k, v := range r.Header {
		log.Println("returned response header")
		fmt.Printf("%v: %s", k, v)
	}
}
