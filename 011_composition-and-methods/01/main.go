package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name string
	Age int
}

func init() {
	tpl = template.Must(template.ParseGlob("index.html"))
}

func main() {

	p1 := person{"Mayk", 28}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", p1)
	if err != nil {
		log.Fatalln(err)
	}

}

