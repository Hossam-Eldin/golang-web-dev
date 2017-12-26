package main

import "time"

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	un           string
	lastActivity time.Time
}
var dbUsers = map[string]user{}       // user ID, user
var dbSessions = map[string]session{} // session ID, session
const sessionLength int = 30
var dbSessionsCleaned time.Time


func main() {

}
