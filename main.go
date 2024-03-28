package main

import (
    "fmt"
    "advocate-cms-2/dao"
	"net/http"
	"html/template"
	"github.com/gorilla/sessions"
)

func main() {
	db, err := dao.DatabaseInit()

	if err != nil {
		fmt.Println("Database initialization failed.")
		return err
	}

	defer db.Close()

	server := &serverCms {
		cookieStore: sessions.NewCookieStore([]byte)),
		db: db,
		mux: http.NewServeMux(),
		&templateMan{}
	}

	fmt.Println("Database initialized")

	// set app routes
	server.mux.HandleFunc(homeEndpoint, server.mainHandler)
	server.mux.HandleFunc(loginEndpoint, server.loginHandler)
	server.mux.HandleFunc(logoutEndpoint, server.logoutHandler)
	server.mux.HandleFunc(uploadImageEndpoint, server.uploadImageHandler)

	fmt.Println("Cms server running on port :8080")
    http.ListenAndServe(":8080", server.mux)
}
