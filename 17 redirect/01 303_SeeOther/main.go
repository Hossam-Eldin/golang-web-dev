package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", one)
	http.HandleFunc("/two", tow)
	http.HandleFunc("/there", there)
	http.ListenAndServe(":8080", nil)
}

func one(w http.ResponseWriter, req *http.Request) {
	fmt.Print("you request method at one:", req.Method, "\n\n")
}

func tow(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar:", req.Method)
	// process form submission here 303
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func there(w http.ResponseWriter, req *http.Request) {
	fmt.Print("you request methods at there", req.Method, "\n\n")
	err := tpl.ExecuteTemplate(w, "index.html", req.Method)
	if err != nil {
		log.Fatalln(err)
	}
}
