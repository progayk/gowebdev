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
	tpl = template.Must(template.ParseGlob("index.gohtml"))
}

func main() {

	p1 := person{"Mayk", 28}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", p1)
	if err != nil {
		log.Fatalln(err)
	}

}

