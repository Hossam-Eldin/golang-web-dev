package main

import (
	"html/template"
	"net/http"
	"fmt"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/lindx", lindx)
	http.HandleFunc("/mindx", mindx)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at inedex:", req.Method, "\n\n")
}

func lindx(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at linedx:", req.Method, "\n\n")

	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
}

func mindx(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at mindx:", req.Method, "\n\n")
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
