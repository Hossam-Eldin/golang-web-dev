package main

import (
	"time"
	"net/http"
	"html/template"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"fmt"
)

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

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/checkUserName", checkUserName)

	http.ListenAndServe(":8080", nil)
}


func index(w http.ResponseWriter, req *http.Request) {
	//show the user using getuser func in session than show session for display and executeTemplate
	u := getUser(w,req)
	showSessions()
	tpl.ExecuteTemplate(w, "index.html", u)
}

func bar(w http.ResponseWriter, req *http.Request)  {
	//show user with getuser func if not logged in send him back to index
	u:=getUser(w,req)
	if !alreadyLoggedIn(w,req){
		http.Redirect(w,req,"/", http.StatusSeeOther)
	}

	//if the role for the user no 007 give him error than execute the bar page

	if u.Role != "007"{
		http.Error(w,"You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w,"bar.html",u)

}



func signup(w http.ResponseWriter, req *http.Request) {
	//show user with getuser func if not logged in send him back to index

	u:=getUser(w,req)
	if !alreadyLoggedIn(w,req){
		http.Redirect(w,req,"/", http.StatusSeeOther)
	}

	var u user
	if req.Method == http.MethodPost {
		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")

		// username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		//create session
		sId := uuid.NewV4()


			c := &http.Cookie{
				Name:  "session",
				Value: sId.String(),
			}
			c.MaxAge =sessionLength
			http.SetCookie(w,c)
			dbSessions[c.Value] = session{un, time.Now()}


		// store user in dbUsers
		//generate password bcrypt
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err !=nil{
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		u= user{un ,bs, f,l,r}
		dbUsers[un] = u
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return

	}
showSessions()
	tpl.ExecuteTemplate(w, "signup.html", nil)
}




func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost{
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return

	}
	showSessions() // for demonstration purposes

	tpl.ExecuteTemplate(w, "login.html", nil)
}





func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	// delete the session
	delete(dbSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	// clean up dbSessions
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}


func checkUserName(w http.ResponseWriter, req *http.Request) {
	sampleUsers := map[string]bool{
		"test@example.com": true,
		"jame@bond.com":    true,
		"moneyp@uk.gov":    true,
	}

	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}

	sbs := string(bs)
	fmt.Println("USERNAME: ", sbs)

	fmt.Fprint(w, sampleUsers[sbs])
}