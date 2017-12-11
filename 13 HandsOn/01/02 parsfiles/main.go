package main

import (
	"net/http"
	"io"
	"html/template"
	"os"
	"log"
)


var tpl *template.Template



func me(w http.ResponseWriter, req *http.Request){
	tpl = template.Must(template.ParseFiles("index.html"))

	err:= tpl.Execute(os.Stdout,"helllo from here")
	if err !=nil{
		log.Fatalln(err)
	}

}
func dog(w http.ResponseWriter, req *http.Request){
	io.WriteString(w,"dog")
}
func index(w http.ResponseWriter, req *http.Request){
	io.WriteString(w,"main")
}


func main() {
	http.HandleFunc("/me/", me)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}
