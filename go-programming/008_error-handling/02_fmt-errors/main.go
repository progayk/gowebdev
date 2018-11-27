package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happened", err)
	}
	defer f.Close()


}
