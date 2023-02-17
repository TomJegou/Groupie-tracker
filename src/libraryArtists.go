package src

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

func removeArtist(artistList []Artist, index int) []Artist {
	artistList[index] = artistList[len(artistList)-1]
	return artistList[:len(artistList)-1]
}

func searchArtists(artistList []Artist, searchContent string) []Artist {
	for indexChar, char := range searchContent {
		for _, artist := range artistList {
			if strings.ToLower(string(artist.Name[indexChar])) != strings.ToLower(string(char)) {
				artistList = removeArtist(artistList, indexChar)
			}
		}
	}
	return artistList
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
		Artists = searchArtists(Artists, searchContent)
		template, errors := template.ParseFiles("static/html/libraryArtists.html")
		if errors != nil {
			fmt.Println("Error Parsing Template")
			fmt.Println(errors)
		}
		template.Execute(w, Artists)
	}
}
