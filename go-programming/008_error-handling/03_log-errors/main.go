package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("no-file.txt")
	if err != nil {
		// log gives a timestamp
		log.Println("err happened", err)
	}
	defer f.Close()
}
