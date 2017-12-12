package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", serve)
	http.HandleFunc("/toby.jpg", img)
	http.ListenAndServe(":8080", nil)
}

func serve(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jpg">`)
}

func img(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "toby.jpg")

}
