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

func sortArtists(sortingOption string) {
	for i := 0; i < len(Artists)-1; i++ {
		x := i
		for j := i + 1; j < len(Artists); j++ {
			if sortingOption == "name" {
				if Artists[j].Name < Artists[x].Name {
					x = j
				}
			} else if sortingOption == "creationDate" {
				if Artists[j].CreationDate < Artists[x].CreationDate {
					x = j
				}
			} else if sortingOption == "numberMembers" {
				if len(Artists[j].Members) < len(Artists[x].Members) {
					x = j
				}
			}
		}
		Artists[i], Artists[x] = Artists[x], Artists[i]
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
		sortArtists("name")
	} else if r.Method == "POST" {
		searchContent := r.FormValue("searchBar")
		if len(searchContent) > 0 {
			searchArtists(searchContent)
		}
	}
	template.Execute(w, Artists)
}
