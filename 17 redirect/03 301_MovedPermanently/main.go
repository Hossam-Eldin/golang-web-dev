package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hos", hos)
	http.HandleFunc("/sam", sam)
	http.ListenAndServe(":8080", nil)
}
func hos(w http.ResponseWriter, req *http.Request) {
	fmt.Print("mothed requested at hos :", req.Method, "\n\n")
}

func sam(w http.ResponseWriter, req *http.Request) {
	fmt.Print("mothed requested at sam :", req.Method, "\n\n")
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}
