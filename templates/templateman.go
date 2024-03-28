package http

import (
	"fmt"
	"html/template"
)

const (
	cmsTmplDir = "http/templates/cms/"
)

type templateMan struct {}

func (t *templateMan) getHomeTmpl() (*Template) {
	homeHtml := cmsTmplDir + "home.html"
	tmpl, _ := template.ParseFiles(homeHtml)
	return tmpl 
}

func (t *templateMan) getLoginTmpl() (*Template) {
	loginHtml := cmsTmplDir + "login.html"
	tmpl, _ := template.ParseFiles(loginHtml)
	return tmpl 
}