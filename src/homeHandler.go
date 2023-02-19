package src

import (
	"fmt"
	"net/http"
	"text/template"
)

func Accueil(w http.ResponseWriter, r *http.Request) {
	OnLibraryArtists = false
	template, errors := template.ParseFiles("static/html/index.html")
	if errors != nil {
		fmt.Println("Error Parsing Template")
		fmt.Println(errors)
	}
	template.Execute(w, nil)
}
