package templates

import (
	"fmt"
	"html/template"
)

const (
	cmsTmplDir = "templates/cms/"
)

type TemplateMan struct {}

func (t *TemplateMan) GetHomeTmpl(w http.ResponseWriter) {
	homeHtml := cmsTmplDir + "home.html"
	tmpl, _ := template.ParseFiles(homeHtml)
	tmpl.Execute(w, nil)
}

func (t *TemplateMan) GetLoginTmpl(w http.ResponseWriter) {
	loginHtml := cmsTmplDir + "login.html"
	tmpl, _ := template.ParseFiles(loginHtml)
	tmpl.Execute(w, nil)
}