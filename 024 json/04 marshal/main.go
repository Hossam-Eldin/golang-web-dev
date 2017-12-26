package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type model struct {
	State    bool
	Pictures []string
}

func main() {
	m :=model{
		State: true,
		Pictures: []string{
			"one.jpg",
			"two.jpg",
			"three.jpg",
		},
	}
	fmt.Println(m)
	bs,err:=json.Marshal(m)
	if err !=nil{
		log.Fatalln(err)
	}
	fmt.Println(string(bs))

}
