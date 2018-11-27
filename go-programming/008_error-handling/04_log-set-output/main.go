package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fl, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer fl.Close()

	// write the logs into log.txt file
	log.SetOutput(fl)

	f, err := os.Open("no-file.txt")
	if err != nil {
		// log gives a timestamp
		log.Println("err happened", err)
	}
	defer f.Close()
}
