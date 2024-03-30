package templates

import (
	"fmt"
	"html/template"
)

const (
	cmsTmplDir = "templates/cms/"
	crudImgTmpl = cmsTmplDir + "crud-image.html"
	imageGalleryTmpl = cmsTmplDir + "image-gallery.html"
	imageGalleryItemTmpl = cmsTmplDir + "image-gallery-item.html"
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

func (t *TemplateMan) GetCrudImgTmpl(w http.ResponseWriter) {
	crudImg := crudImage {
		Title: updateImageTitle,
		Endpoint: updateImageTitle,
	}
}