package http

import (
    "fmt"
    "advocate-cms-2/internal/dao"
	"advocate-cms-2/internal/templates"
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

type ServerCms struct {
	cookieStore *CookieStore
	db *dao.Dao
	mux *ServeMux
	*templates.Templates
}

func ServerInit() (*ServerCms) error {

	db, err := dao.DatabaseInit()

	if err != nil {
		fmt.Println("Database initialization failed.")
		return err
	}

	return &ServerCms {
		cookieStore: sessions.NewCookieStore([]byte)),
		db: db,
		mux: http.NewServeMux(),
		&template.Templates{}
	}, nil
}

func (server *ServerCms) ServerShutdown() {
	if server.db != nil {
		server.db.Close()
	}
}

func (server *ServerCms) ServerStart() {
	// set app routes
	server.mux.HandleFunc(homeEndpoint, server.mainHandler)
	server.mux.HandleFunc(loginEndpoint, server.loginHandler)
	server.mux.HandleFunc(logoutEndpoint, server.logoutHandler)
	server.mux.HandleFunc(uploadImageEndpoint, server.uploadImageHandler)

	fmt.Println("Cms server running on port :8080")
    http.ListenAndServe(":8080", server.mux)
}

func (server *ServerCms) authenticateUser(w http.ResponseWriter, r *http.Request) {
	// is user authenticated
	session, _ := server.cookieStore.Get(r, sessionKey)
	_, ok := session.Values[sessionId].(string)

	// not authenticated redire to login
	if !ok{
		http.Redirect(w, r, loginEndpoint, http.StatusSeeOther)
	}
}

func (server *ServerCms) mainHandler(w http.ResponseWriter, r *http.Request) {

	// only serving get methods from here
	if r.Method != http.MethodGet {
		http.Error(w, httpErrorMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	// redirects to login if not authenticated
	server.authenticateUser(w, r)

	// else return home page
	server.GetCmsHome().Execute(w, nil)
}