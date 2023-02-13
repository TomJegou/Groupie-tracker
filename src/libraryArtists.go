package src

import (
	"net/http"
	"text/template"
)

func libraryArtists(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("static/html/libraryArtists.html")
	template.Execute(w, Artists)
}
