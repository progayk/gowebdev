package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

func init() {
	tpl = template.Must(template.ParseGlob("index.gohtml"))
}

func main() {

	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{"CSCI-40", "Introduction to Programming in Go", "4"},
				{"CSCI-130", "Introduction to Web Programming with Go", "4"},
				{"CSCI-140", "Mobile Apps Using Go", "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{"CSCI-50", "Advanced Go", "5"},
				{"CSCI-190", "Advanced Web Programming with Go", "5"},
				{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
			},
		},
		Summer: semester{
			Term: "Summer",
			Courses: []course{
				{"CSCI-60", "Ultimate Go", "8"},
				{"CSCI-290", "Ultimate Web Programming with Go", "2"},
			},
		},
	}


	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", y)
	if err != nil {
		log.Fatalln(err)
	}

}

