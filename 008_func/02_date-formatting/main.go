package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func FormatDateDMY(t time.Time) string {
	return t.Format("02_print-uri Jan 06 15:04 MST")
}

func Kitchen(t time.Time) string {
	return t.Format(time.Kitchen)
}

var fm = template.FuncMap{
	"fmtDateDMY": FormatDateDMY,
	"Kitchen": Kitchen,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}

}

