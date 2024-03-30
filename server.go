package main

import (
    "fmt"
    "advocate-cms-2/dao"
	"advocate-cms-2/templates"
	"net/http"
	"html/template"
	"github.com/gorilla/sessions"
)

const (
	devSecretKey = "decSecretKey" // for dev purposes only
	homeEndpoint = "/"
	loginEndpoint = "/login"
	logoutEndpoint = "/logout"
	uploadImageEndpoint = "/upload-image"
    httpErrorMethodNotAllowed = "Method Not Allowed"
    httpErrorInternalError = "Internal Server Error"
    httpErrorBadLogin = "Invalid Login"
	sessionKey = "session-name"
	sessionId = "username"
)

type serverCms struct {
	cookieStore *CookieStore
	db *dao.Dao
	mux *ServeMux
	*templates.TemplateMan
}

func (server *ServerCms) authenticateUser(w http.ResponseWriter, r *http.Request) {
	// is user authenticated
	session, _ := server.cookieStore.Get(r, sessionKey)
	_, ok := session.Values[sessionId].(string)

	// not authenticated redirect to login
	if !ok{
		http.Redirect(w, r, loginEndpoint, http.StatusSeeOther)
	}
}

func (server *ServerCms) mainHandler(w http.ResponseWriter, r *http.Request) {

	// only serving get methods
	if r.Method != http.MethodGet {
		http.Error(w, httpErrorMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	// redirects to login if not authenticated
	server.authenticateUser(w, r)

	// else return home page
	server.GetHomeTmpl(w)
}

func (server *ServerCms) loginHandler(w http.ResponseWriter, r *http.Request) {

	// only post and get allowed
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, httpErrorMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	// get login
	if r.Method == http.MethodGet {
		server.GetLoginTmpl(w)
		return
	}

	//post
	// get credentials from body post to validate login
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	loginStatus, err := server.db.IsValidLogin(username, password)

	if err != nil {
		fmt.Println("Error validating login.")
		http.Error(w, httpErrorInternalError, http.StatusInternalServerError)
		return
	}
	
	if !loginStatus {
		http.Error(w, httpErrorBadLogin, http.StatusBadRequest)
		return
	}

	// create session and redirect to home
	session, _ := server.cookieStore.Get(r, sessionKey)
	session.Values[sessionId] = username
	session.Save(r, w)

	http.Redirect(w, r, homeEndpoint, http.StatusSeeOther)
}

func (server *ServerCms) logoutHandler(w http.ResponseWriter, r *http.Request) {
	// post only
	if r.Method != http.MethodPost {
		http.Error(w, httpErrorMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	// redirects to login if not authenticated
	server.authenticateUser(w, r)

	// delete session
    session, _ := server.cookieStore.Get(r, sessionKey)
    delete(session.Values, sessionId)
    session.Save(r, w)

    // Redirect to the login page.
    http.Redirect(w, r, loginEndpoint, http.StatusSeeOther)
}

func (server *ServerCms) uploadImageHandler(w http.ResponseWriter, r *http.Request) {
	
	// only post and get allowed
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, httpErrorMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	// redirects to login if not authenticated
	server.authenticateUser(w, r)

	if r.Method == http.MethodGet {
		server
	}
}