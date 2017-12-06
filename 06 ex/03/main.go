package main

import (
	"html/template"
	"os"
	"log"
)

type Breakfast struct {
	items []string
}

type Lunch struct {
	items []string
}
type Dinner struct {
	items []string
}

type restaurant struct {
	Breakfast Breakfast
	Dinner    Dinner
	Lunch     Lunch
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {

	meals := restaurant{
		Breakfast{
			items: []string{"food1", "food2"},
		},
		Dinner{
			items: []string{"food1", "food2"},
		},
		Lunch{
			items: []string{"food1", "food2"},
		},
	}

	err := tpl.Execute(os.Stdout,meals)
	if err != nil {
		log.Fatalln(err)
	}
}
