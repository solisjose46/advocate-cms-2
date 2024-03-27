package templates

import (
	"fmt"
	"html/template"
)

type Templates struct {}

func (t *Templates) GetHome() (*Template) {
	homeHtml := "frontend/html/home.html"
	tmpl, _ := template.ParseFiles(homeHtml)
	return tmpl 
}

func (t *Templates) GetCmsHome() (*Template) {
	homeHtml := "frontend/html/home.html"
	tmpl, _ := template.ParseFiles(homeHtml)
	return tmpl 
}