package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", link)
	http.ListenAndServe(":8080", nil)
}

func link(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="GET">
	 <input type="text" name="q">
	 <input type="submit">
	</form>
	<br>`+v)
}
