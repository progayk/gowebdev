package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type Page struct {
	Title, Heading, Input string
}

func init() {
	tpl = template.Must(template.ParseGlob("index.html"))
}

func main() {

	p := Page{
		Title: "Not escaped",
		Heading: "this is text/template and it's not escaped.",
		Input: "<script>alert('yey yey yey');</script>",
	}


	err := tpl.ExecuteTemplate(os.Stdout, "index.html", p)
	if err != nil {
		log.Fatalln(err)
	}

}

