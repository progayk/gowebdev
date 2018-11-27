package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Person struct {
	First string
	Last string
}

func (p *Person) speak(s *string) {
	fmt.Printf("Hello! %s says %s\n", p.First, *s)
}

func main() {
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])


	p1 := Person{
		First: os.Args[1],
		Last: os.Args[2],
	}

	p1.speak(&os.Args[3])

	tpl := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	    <meta charset="UTF-8">
	    <title>Hello World</title>
	</head>
	<body>
		<h1>` + p1.First + p1.Last +`</h1>
	</body>
	</html>
	`)


	nf, err := os.Create("template.gohtml")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	xr := strings.NewReader(tpl)
	fmt.Println("The length of tpl string:", xr.Len())

	if _, err = io.Copy(nf, strings.NewReader(tpl)); err != nil {
		log.Fatal("the file couldnt be written", err)
	}

}
