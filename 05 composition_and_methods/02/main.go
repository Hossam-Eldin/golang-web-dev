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

type doubleZero struct {
	Person
	LincesToKill bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	p := doubleZero{
		Person{
			Name: "Ian Fleming",
			Age:  56,
		}, false,
	}

	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}
}
