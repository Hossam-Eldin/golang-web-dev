package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("Session")
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "Session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
