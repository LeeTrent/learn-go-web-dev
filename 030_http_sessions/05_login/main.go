package main

import (
	"html/template"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

const cookieName string = "session"

var userMgr UserMgr
var sessionMgr SessionMgr
var tpl *template.Template

func init() {
	userMgr = UserMgr{}
	sessionMgr = SessionMgr{}
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	user := getUser(w, req)
	tpl.ExecuteTemplate(w, "index.gohtml", user)
}

func bar(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	user := getUser(w, req)
	tpl.ExecuteTemplate(w, "bar.gohtml", user)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		un := req.FormValue("username")
		pw := req.FormValue("password")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")

		// username taken?
		if userMgr.userNameIsTaken(un) {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// encrypt password
		encyptedPW, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// create user
		user, err := userMgr.createUser(encyptedPW, un, fn, ln)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// create session
		session := sessionMgr.createSession(user.UserName)

		// create and set cookie
		CreateAndSetCookie(cookieName, session.sessionId, w)

		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func getUser(w http.ResponseWriter, req *http.Request) User {
	cookie := GetAndSetCookie(cookieName, w, req)
	sessionId := cookie.Value
	userName := sessionMgr.getUserName(sessionId)
	user := userMgr.getUser(userName)
	return user
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	if HasCookie(cookieName, req) == false {
		return false
	}
	cookie := GetAndSetCookie(cookieName, w, req)
	sessionId := cookie.Value
	userName := sessionMgr.getUserName(sessionId)
	_, ok := userMgr[userName]
	return ok
}
