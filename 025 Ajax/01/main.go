package main

import (
	"html/template"
	"net/http"
	"fmt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/joo",joo)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request)  {
	tpl.ExecuteTemplate(w,"index.html",nil)
}

func joo(w http.ResponseWriter, req *http.Request)  {
	s:="this from jojo"
	fmt.Fprintln(w,s)
}
