package src

import (
	"net/http"
	"text/template"
)

func AboutHandlerFunc(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("static/html/about.html")
	template.Execute(w, nil)
}
