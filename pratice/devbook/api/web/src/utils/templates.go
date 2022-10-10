package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func LoadTemplate() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func ExecuteTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
	// if erro != nil {
	// 	log.Fatal(erro)
	// }
}
