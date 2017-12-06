package main

import (
	"html/template"
	"os"
	"log"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("time.html"))
}

func monthDayYear(t time.Time ) string {
		return  t.Format("01-02-2006")
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {


	err := tpl.ExecuteTemplate(os.Stdout, "time.html" ,time.Now())
	if err != nil{
		log.Fatalln(err)
	}

}
