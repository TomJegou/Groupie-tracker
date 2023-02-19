package src

import (
	"fmt"
	"net/http"
	"text/template"
)

func AboutHandlerFunc(w http.ResponseWriter, r *http.Request) {
	OnLibraryArtists = false
	template, errors := template.ParseFiles("static/html/about.html")
	if errors != nil {
		fmt.Println("Error Parsing Template")
		fmt.Println(errors)
	}
	template.Execute(w, nil)
}
