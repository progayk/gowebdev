package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	data := 42

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", data)
	if err != nil {
		log.Fatalln(err)
	}

}

