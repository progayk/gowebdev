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

func (p person) SomeProcessing() int {
	return 7
}

func (p person) TakesArgument(x int) int {
	return x * 2
}

func (p person) DblAge() int {
	return p.Age * 2
}

func init() {
	tpl = template.Must(template.ParseGlob("index.gohtml"))
}

func main() {

	p1 := person{
		Name: "Jony",
		Age: 28,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", p1)
	if err != nil {
		log.Fatalln(err)
	}

	// Generally speaking, best practice:
	// call functions in templates for formatting only; not processing or ...

	// The main reasons that you don't want to do any data processing in your templates:
	// (1) seperation of concerns
	// (2) if you're using a function more than once in your template
	// the server needs to do the processing more than once.
	// (though the standard library might cache processing.
	// yet I don't know how. If you find it let me know ;) )

}

