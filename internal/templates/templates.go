package templates

import (
	"fmt"
	"html/template"
)

type CrudImage struct {
	Title string
	Endpoint string
}

type Templates struct {

}

func (t *Templates) GetCmsHome() (*Template) {
	homeHtml := "frontend/html/home.html"
	tmpl, _ := template.ParseFiles(homeHtml)
	return tmpl 
}