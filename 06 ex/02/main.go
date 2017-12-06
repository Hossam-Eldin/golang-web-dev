package main

import (
	"html/template"
	"os"
	"log"
)

var tpl *template.Template



type cHotel  struct {
	Name string
	Address string
	City string
	Zip string
	Region string
}

type hotels []cHotel

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {

	h := hotels{
		cHotel{
			Name:    "Hotel California",
			Address: "42 Sunset Boulevard",
			City:    "Los Angeles",
			Zip:     "95612",
			Region:  "southern",
		},
		cHotel{
			Name:    "H",
			Address: "4",
			City:    "L",
			Zip:     "95612",
			Region:  "southern",
		},
	}

	err :=tpl.Execute(os.Stdout,h)
	if err != nil{
		log.Fatalln(err)
	}
}