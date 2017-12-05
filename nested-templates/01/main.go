package main

import (
	"html/template"
	"os"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*.html"))
}

func main() {

	err:= tpl.ExecuteTemplate(os.Stdout,"index.html" ,42)

	if err !=nil{
		log.Fatalln(err)
	}

}
