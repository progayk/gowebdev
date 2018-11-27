package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Printf("%T\n", *resp)
	fmt.Println(*resp)
}
