package main

import (
	"html/template"
	"log"
	"os"
)

type item struct {
	Name, Descrip string
	Price         float64
}

type meal struct {
	Meal string
	Item []item
}

type menu []meal

type restaurant struct {
	Name string
	Menu menu
}

type restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {

	r := restaurants{
		restaurant{
			Name: "Kitchen",
			Menu: menu{
				meal{
					Meal: "BreakFast",
					Item: []item{
						item{
							Name:    "Oatmeal",
							Descrip: "yum yum",
							Price:   4.95,
						},
						item{
							Name:    "Cheerios",
							Descrip: "American eating food traditional now",
							Price:   3.95,
						},
						item{
							Name:    "Juice Orange",
							Descrip: "Delicious drinking in throat squeezed fresh",
							Price:   2.95,
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}
