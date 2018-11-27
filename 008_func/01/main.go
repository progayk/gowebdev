package main

import (
	"log"
	"os"
	"strings"
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

// create a FuncMap to register functions.
// "uc" is what the func will be called in the template.
// "uc" is the ToUpper func from package strings
// "ft" is the func I declared
// "ft" slices a string, return the first three chars.
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": FirstThree,
}

func FirstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
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

	data := struct{
		Wisdom []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}

