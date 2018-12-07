package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {

	u1 := user{
		Name:"Buddha",
		Motto:"to believe or not to believe",
		Admin: false,
	}

	u2 := user{
		Name:"Jesus",
		Motto:"What is love",
		Admin: true,
	}

	u3 := user{
		Name:"",
		Motto:"nothing",
		Admin: true,
	}

	users := []user{u1,u2,u3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}

}

