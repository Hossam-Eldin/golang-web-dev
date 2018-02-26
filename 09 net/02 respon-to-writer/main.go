package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("my-key", "this a key")
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>function here</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
