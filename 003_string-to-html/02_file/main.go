package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Mayk Jony"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	    <meta charset="UTF-8">
	    <title>Hello World</title>
	</head>
	<body>
		<h1>` + name + `</h1>
	</body>
	</html>
	`

	nf, err := os.Create("template.gohtml")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	xr := strings.NewReader(tpl)
	fmt.Println("The length of tpl string:", xr.Len())

	io.Copy(nf, strings.NewReader(tpl))

}
