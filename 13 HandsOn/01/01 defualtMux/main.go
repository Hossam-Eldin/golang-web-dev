package main

import (
	"io"
	"net/http"
)

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "me")
}
func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog")
}
func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "main")
}

func main() {
	http.HandleFunc("/me/", me)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}
