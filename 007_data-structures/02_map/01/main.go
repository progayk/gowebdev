package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	sages := map[string]string{
			"India": "Gandhi",
			"America": "MLK",
			"Meditate": "Buddha",
			"Love": "Jesus",
			"Prohpet": "Muhammad",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}

}

