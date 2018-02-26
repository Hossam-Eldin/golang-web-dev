package main

import (
	"io"
	"net/http"
)

func serve(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/public/toby.jpg">`)

}

func main() {
	http.HandleFunc("/index", serve)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}
