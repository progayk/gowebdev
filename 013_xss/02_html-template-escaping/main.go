package main

import (
	"log"
	"os"
	"html/template"
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
		Title: "escaped ",
		Heading: "this is html/template and it's escaped.",
		Input: "<script>alert('yey yey yey');</script>",
	}


	err := tpl.ExecuteTemplate(os.Stdout, "index.html", p)
	if err != nil {
		log.Fatalln(err)
	}

}

