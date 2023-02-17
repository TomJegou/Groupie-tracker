package src

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

func searchArtists(artistList []Artist, artistListResult []Artist, searchContent string) {
	for _, artist := range artistList {
		isOk := true
		for indexChar, char := range searchContent {
			if !strings.EqualFold(string(artist.Name[indexChar]), string(char)) {
				isOk = false
				break
			}
		}
		if isOk {
			artistListResult = append(artistListResult, artist)
		}
	}
}

func libraryArtists(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		template, errors := template.ParseFiles("static/html/libraryArtists.html")
		if errors != nil {
			fmt.Println("Error Parsing Template")
			fmt.Println(errors)
		}
		template.Execute(w, Artists)
	} else if r.Method == "POST" {
		searchContent := r.FormValue("searchBar")
		searchArtistsListResult := []Artist{}
		searchArtists(Artists, searchArtistsListResult, searchContent)
		template, errors := template.ParseFiles("static/html/libraryArtists.html")
		if errors != nil {
			fmt.Println("Error Parsing Template")
			fmt.Println(errors)
		}
		template.Execute(w, searchArtistsListResult)
	}
}
