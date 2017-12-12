package main

import (
	"net/http"
	"html/template"
	"log"
)

var tpl *template.Template

func init() {
	tpl=template.Must(template.ParseGlob("assets/*"))

}
func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/about",about)
	http.ListenAndServe(":8080", nil)
}


func index(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w,"index.html",nil)
	if err !=nil{
		log.Fatalln(err)
	}
}


func about(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w,"about.html",nil)
	if err !=nil{
		log.Fatalln(err)
	}
}



func handleErr(w http.ResponseWriter,err error){
	if err !=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

