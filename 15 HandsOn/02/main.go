package main

import (
	"net/http"
	"html/template"
	"log"
)

var tpl *template.Template

func main() {
	f:= http.FileServer(http.Dir("."))

	http.HandleFunc("/", index)
	http.Handle("/all",f)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl = template.Must(template.ParseFiles("index.html"))
	err := tpl.Execute(w, "hello")
	if err != nil {
		log.Fatalln(err)
	}
}

func serve (w http.ResponseWriter, req *http.Request){
}
