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

type car struct {
	Manufacturer string
	Model string
	Doors int
}

type items struct {
	Wisdom []sage
	Transport []car
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

	corolla := car{
		Manufacturer: "Toyota",
		Model: "Corolla",
		Doors: 4,
	}

	f150 := car{
		Manufacturer: "Ford",
		Model: "F150",
		Doors: 2,
	}

	sages := []sage{mlk, jesus, buddha, muhammad}
	cars := []car{corolla, f150}

	data := items{
		Wisdom: sages,
		Transport:cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

}

