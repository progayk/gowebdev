package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
)

var nodeCode = `
	function squareIt(val) {
		return val * val
	}
	let a = 12
	console.log(a, 'squared is', squareIt(a));
`

func main() {
	// create an temporary file with .js extension
	tempFile, err := ioutil.TempFile("", "*.js")
	if err != nil {
		log.Fatalln(err)
	}
	// write the node code to it
	io.WriteString(tempFile, nodeCode)
	tempFile.Close()

	// run node with file as argument
	cmd := exec.Command("node", tempFile.Name())

	b, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))
}
