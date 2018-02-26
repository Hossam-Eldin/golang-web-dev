package main

import (
	"html/template"
	"log"
	"os"
)

type Person struct {
	Name string
	Age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("01 template/*.html"))
}

func main() {

	p := Person{
		Name: "Archer",
		Age:  35,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", p)

	if err != nil {
		log.Fatalln(err)
	}

}
