package main

import (
	"html/template"
	"os"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("map.html"))
}

func main() {
	sages := map[string]string{
		"india"	    :"Gandhi",
		"usa"	    :"MLK",
		"monk"      : "Buddha",
		"in the sky": "Jesus",
		"mercy" : "Muhammad"}
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}

}
