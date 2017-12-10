package main

import (
	"net/http"
	"io"
	"github.com/gorilla/mux"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "who let the dogs out ??! ")
}

type hotcat int

func (c hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "cat cat cat")
}


func main() {
	var d hotdog
	var c hotcat

	http.Handle("/dog/", d)
	http.Handle("/cat", c)

	http.ListenAndServe(":8080", nil)
}
