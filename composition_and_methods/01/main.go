package main

import (
	"html/template"
	"os"
	"log"
)

type Person struct {
	Name string
	Age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*.html"))
}

func main() {


	p :=Person{
		Name: "Archer",
		Age: 35,
	}
	err:= tpl.ExecuteTemplate(os.Stdout,"index.html" ,p)

	if err !=nil{
		log.Fatalln(err)
	}

}
