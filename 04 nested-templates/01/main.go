package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("01 template/*.html"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", 42)

	if err != nil {
		log.Fatalln(err)
	}

}
