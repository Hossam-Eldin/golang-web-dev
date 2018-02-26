package main

import (
	"github.com/satori/go.uuid"
	"net/http"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	//set cookie
	c, err := req.Cookie("session")
	if err != nil {
		sId := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sId.String(),
		}
	}
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u

}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
