package main

import (
	"net/http"
	"html/template"
	"log"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/", index)
		http.HandleFunc("/deadpool.jpg", serve)
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
	http.ServeFile(w, req, "deadpool.jpg")
}
