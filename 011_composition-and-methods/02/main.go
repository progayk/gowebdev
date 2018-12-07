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

type doubleZero struct {
	person
	LicenseToKill bool
}

func init() {
	tpl = template.Must(template.ParseGlob("index.html"))
}

func main() {

	dz := doubleZero{
		 person: person{
		 	"Mayk",
		 	28,
		},
		 LicenseToKill: false,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", dz)
	if err != nil {
		log.Fatalln(err)
	}

}

