package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

func init() {
	tpl= template.Must(template.ParseFiles("index.html"))
}

func main() {
	err := tpl.Execute(os.Stdout,42)
	if err !=nil{
		log.Fatalln(err)
	}

}
