package main

import (
    "fmt"
    "advocate-cms-2/internal/dao"
	"net/http"
	"html/template"
	"github.com/gorilla/sessions"
)

const (
	homeEndpoint = "/"
	loginEndpoint = "/login"
	logoutEndpoint = "/logout"
    httpErrorMethodNotAllowed = "Method Not Allowed"
    httpErrorInternalError = "Internal Server Error"
    httpErrorBadLogin = "Invalid Login"
	sessionKey = "session-name"
	sessionId = "username"
)

// this key for dev only obviously
var store = sessions.NewCookieStore([]byte("devSecretKey"))

var db *dao.Dao

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// check if user is authenticated
	session, _ := store.Get(r, sessionKey)

	_, ok := session.Values[sessionId].(string)

	// if not authenticated redirect to /login
	if !ok {
		http.Redirect(w, r, loginEndpoint, http.StatusSeeOther)
	}

	// else return home page
    homeHtml := "frontend/html/home.html"
	templ, _ := template.ParseFiles(homeHtml)
	templ.Execute(w, nil)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {

	// post only
	if r.Method != http.MethodPost {
		http.Error(w, httpErrorMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

    session, _ := store.Get(r, sessionKey)
    delete(session.Values, sessionId)
    session.Save(r, w)

    // Redirect to the login page.
    http.Redirect(w, r, loginEndpoint, http.StatusSeeOther)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	// get login page
    loginHtml := "frontend/html/login.html"
	if r.Method == http.MethodGet {
		templ, err := template.ParseFiles(loginHtml)
        if err != nil {
            http.Error(w, httpErrorInternalError, http.StatusInternalServerError)
            return
        }

		templ.Execute(w, nil)
		return
	}

	// only post allowed after this
	if r.Method != http.MethodPost {
		http.Error(w, httpErrorMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	// get credentials from body post
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	// validate login
	loginStatus, err := db.IsValidLogin(username, password)

	if err != nil {
		http.Error(w, httpErrorInternalError, http.StatusInternalServerError)
		return
	}
	
	if !loginStatus {
		http.Error(w, httpErrorBadLogin, http.StatusBadRequest)
		return
	}

	// create session and redirect to home
	session, _ := store.Get(r, sessionKey)
	session.Values[sessionId] = username
	session.Save(r, w)

	http.Redirect(w, r, homeEndpoint, http.StatusSeeOther)
}

func main() {
    var err error
    db, err = dao.DatabaseInit()

    if err != nil {
        fmt.Println("Database initialization failed.")
        return
    }

    defer db.Cms.Close()
    defer db.Advocate.Close()

	fmt.Println("Database initialized!")

	// default server mux that is NOT global
	mux := http.NewServeMux()

	// set app routes
	mux.HandleFunc(homeEndpoint, mainHandler)
	mux.HandleFunc(loginEndpoint, loginHandler)
	mux.HandleFunc(logoutEndpoint, logoutHandler)


	fmt.Println("Cms server running on port :8080")
    http.ListenAndServe(":8080", mux)
}
