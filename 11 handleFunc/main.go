package main

import (
	"net/http"
	"io"
)

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "cat")
}

func main() {
	http.HandleFunc("/cat", c)
	http.ListenAndServe(":8080", nil)
}
