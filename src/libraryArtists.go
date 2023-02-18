package src

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

func setArtistVisibility(a *Artist, isVisible bool) {
	a.IsVisible = isVisible
}

func setAllArtistVisibility(isVisible bool) {
	for i := 0; i < len(Artists); i++ {
		setArtistVisibility(&Artists[i], isVisible)
	}
}

func searchArtists(searchContent string) {
	setAllArtistVisibility(false)
	for i := 0; i < len(Artists); i++ {
		isOk := true
		for indexChar, char := range searchContent {
			if !strings.EqualFold(string(Artists[i].Name[indexChar]), string(char)) {
				isOk = false
				break
			}
		}
		if isOk {
			setArtistVisibility(&Artists[i], true)
		}
	}
}

func libraryArtists(w http.ResponseWriter, r *http.Request) {
	PutBodyResponseApiIntoStruct(URLARTISTS, &Artists)
	template, errors := template.ParseFiles("static/html/libraryArtists.html")
	if errors != nil {
		fmt.Println("Error Parsing Template")
		fmt.Println(errors)
	}
	if r.Method == "GET" {
		setAllArtistVisibility(true)
		template.Execute(w, Artists)
	} else if r.Method == "POST" {
		searchContent := r.FormValue("searchBar")
		searchArtists(searchContent)
	}
	template.Execute(w, Artists)
}
