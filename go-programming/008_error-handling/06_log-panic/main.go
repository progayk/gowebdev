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
		log.Panicln(err)
	}
}

func foo() {
	fmt.Println("When the panic() is called deferred funcs are run.")
}
