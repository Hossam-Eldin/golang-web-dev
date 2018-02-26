package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/", serve)
	http.Handle("/public", http.StripPrefix("/public", http.FileServer(http.Dir("/assets"))))
	http.ListenAndServe(":8080", nil)
}

func serve(w http.ResponseWriter, req *http.Request) {
	tpl = template.Must(template.ParseGlob("assets/*"))
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
