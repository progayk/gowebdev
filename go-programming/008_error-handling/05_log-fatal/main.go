package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		// Fatal functions call os.Exit(1) after writing log message
		log.Fatalln(err)
	}
}

func foo() {
	fmt.Println("When the os.Exit(1) is called deferred funcs are not run.")
}
