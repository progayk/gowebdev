package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {

	g1 := struct {
		Score1 int
		Score2 int
	}{
		7,
		9,
	}

	err := tpl.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatalln(err)
	}

}

