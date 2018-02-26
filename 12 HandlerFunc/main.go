package main

import (
	"io"
	"net/http"
)

func c(w http.ResponseWriter, req http.Request) {
	io.WriteString(w, "cats ")
}

func main() {
	http.Handle("/cat", http.HandlerFunc(c))
	http.ListenAndServe(":8080", nil)
}
