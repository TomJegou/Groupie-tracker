package src

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

func searchArtists(artistList []Artist, artistListResult []Artist, searchContent string) {
	for indexChar, char := range searchContent {
		for _, artist := range artistList {
			if strings.EqualFold(string(artist.Name[indexChar]), strings.ToLower(string(char))) {
				artistListResult = append(artistListResult, artist)
			}
		}
	}
}

func libraryArtists(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
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
