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

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", buddha)
	if err != nil {
		log.Fatalln(err)
	}

}

