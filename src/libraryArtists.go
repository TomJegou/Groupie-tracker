package src

import (
	"fmt"
	"net/http"
	"text/template"
)

func libraryArtists(w http.ResponseWriter, r *http.Request) {
	template, errors := template.ParseFiles("static/html/libraryArtists.html")
	if errors != nil {
		fmt.Println("Error Parsing Template")
		fmt.Println(errors)
	}
	template.Execute(w, Artists)
}
