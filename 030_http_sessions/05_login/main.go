package main

import (
	"github.com/LeeTrent/cookieutil"
	"github.com/LeeTrent/sessionmgr"
	"github.com/LeeTrent/usermgr"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

const cookieName string = "session"

var userMgr *usermgr.UserMgr
var sessionMgr *sessionmgr.SessionMgr
var tpl *template.Template

func init() {
	userMgr = usermgr.NewUserMgr()
	sessionMgr = sessionmgr.NewSessionMgr()
	tpl = template.Must(template.ParseGlob("templates/*"))

	// seed database with at least one user
	bs, _ := bcrypt.GenerateFromPassword([]byte("mypassword"), bcrypt.MinCost)
	userMgr.CreateUser(bs, "casey@casey.com", "Casey", "Trent")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
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
		if userMgr.UserNameIsTaken(un) {
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
		user, err := userMgr.CreateUser(encyptedPW, un, fn, ln)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// create session
		session := sessionMgr.CreateSession(user.UserName)

		// create and set cookie
		cookieutil.CreateAndSetCookie(cookieName, session.SessionId, w)

		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func login(rw http.ResponseWriter, req *http.Request) {

	if alreadyLoggedIn(rw, req) {
		http.Redirect(rw, req, "/", http.StatusSeeOther)
		return
	}

	// process form login form submission
	if req.Method == http.MethodPost {
		// get form values
		un := req.FormValue("username")
		pw := req.FormValue("password")


		// validate user name
		user, found := userMgr.GetUser(un)
		if found == false {
			http.Error(rw, "Invalid user name or password", http.StatusForbidden)
			return
		}

		// validate password
		err := bcrypt.CompareHashAndPassword(user.Password, []byte(pw))
		if err != nil {
			http.Error(rw, "Invalid user name or password", http.StatusForbidden)
			return
		}

		// create session
		session := sessionMgr.CreateSession(user.UserName)

		// create and set cookie
		cookieutil.CreateAndSetCookie(cookieName, session.SessionId, rw)

		// redirect
		http.Redirect(rw, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(rw, "login.gohtml", nil)
}

func getUser(w http.ResponseWriter, req *http.Request) usermgr.User {
	cookie := cookieutil.GetAndSetCookie(cookieName, w, req)
	sessionId := cookie.Value
	userName := sessionMgr.GetUserName(sessionId)
	user, _ := userMgr.GetUser(userName)
	return user
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	if cookieutil.HasCookie(cookieName, req) == false {
		return false
	}
	cookie := cookieutil.GetAndSetCookie(cookieName, w, req)
	sessionId := cookie.Value
	userName := sessionMgr.GetUserName(sessionId)
	_, ok := userMgr.GetUser(userName)
	return ok
}
