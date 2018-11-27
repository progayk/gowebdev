package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	buddha := sage{
		Name: "Buddha",
		Motto: "The belief of no beliefs",
	}

	muhammad := sage{
		Name: "Muhammad",
		Motto: "Allah bir",
	}

	mlk := sage{
		Name: "Martin Luther King",
		Motto: "hatred never ceased with hatred but love alone is healed",
	}

	jesus := sage{
		Name: "Jesus",
		Motto: "What is love? Baby don't hurt me, no more!",
	}

	sages := []sage{mlk, jesus, buddha, muhammad}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}

}

