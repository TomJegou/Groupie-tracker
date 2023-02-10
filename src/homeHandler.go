package src

import (
	"net/http"
	"text/template"
)

func Accueil(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("static/html/index.html")
	template.Execute(w, nil)
}
