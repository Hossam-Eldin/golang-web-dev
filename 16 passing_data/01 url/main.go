package main

import (
	"net/http"
	"fmt"
)

func main() {
http.HandleFunc("/", url)
http.Handle("/fiv.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}


func url(w http.ResponseWriter, req *http.Request){
	v:=req.FormValue("q")
	fmt.Fprintln(w, "do my search:" +v)
}