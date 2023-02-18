package src

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

func makeAllArtistVisible(truc *Artist) {
	truc.IsVisible = true
}

func searchArtists(artistList []Artist, searchContent string) []Artist {
	result := []Artist{}
	for _, artist := range artistList {
		isOk := true
		for indexChar, char := range searchContent {
			if !strings.EqualFold(string(artist.Name[indexChar]), string(char)) {
				isOk = false
				break
			}
		}
		if isOk {
			result = append(result, artist)
		}
	}
	return result
}

func libraryArtists(w http.ResponseWriter, r *http.Request) {
	PutBodyResponseApiIntoStruct(URLARTISTS, &Artists)
	if r.Method == "GET" {
		template, errors := template.ParseFiles("static/html/libraryArtists.html")
		if errors != nil {
			fmt.Println("Error Parsing Template")
			fmt.Println(errors)
		}
		template.Execute(w, Artists)
	} else if r.Method == "POST" {
		searchContent := r.FormValue("searchBar")
		searchArtistsListResult := searchArtists(Artists, searchContent)
		template, errors := template.ParseFiles("static/html/libraryArtists.html")
		if errors != nil {
			fmt.Println("Error Parsing Template")
			fmt.Println(errors)
		}
		template.Execute(w, searchArtistsListResult)
	}
}
