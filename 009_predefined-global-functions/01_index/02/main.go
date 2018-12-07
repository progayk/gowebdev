package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {

	xs := []string{"zero", "one", "two", "three", "four"}

	data := struct {
		Words []string
		Lname string
	}{
		Words:xs,
		Lname:"Jony",
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

}

