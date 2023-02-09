package src

import (
	"net/http"
	"text/template"
)

func ArtistsDetailsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("static/html/artistsDetails.html")
	template.Execute(w, nil)
}
