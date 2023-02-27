package src

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

type LibraryArtists struct {
	Artistlist    *[]Artist
	SortingFilter string
	Asc           bool
}

var LibArtists LibraryArtists

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

func reverseSliceArtist() {
	for i := 0; i < len(Artists)/2; i++ {
		Artists[i], Artists[len(Artists)-1-i] = Artists[len(Artists)-1-i], Artists[i]
	}
}

func selectionSort(){

}

func sortArtists(sortingOption string, asc bool) {
	for i := 0; i < len(Artists)-1; i++ {
		x := i
		for j := i + 1; j < len(Artists); j++ {
			if sortingOption == "name" {
				if strings.ToLower(Artists[j].Name) < strings.ToLower(Artists[x].Name) {
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
			} else if sortingOption == "Firstalbumrelease" {
				SortFirstAlbum()
			}
		}
		Artists[i], Artists[x] = Artists[x], Artists[i]
	}
	if !asc {
		reverseSliceArtist()
	}
}

func libraryArtists(w http.ResponseWriter, r *http.Request) {
	if !OnLibraryArtists {
		PutBodyResponseApiIntoStruct(URLARTISTS, &Artists)
		OnLibraryArtists = true
	}
	LibArtists.Artistlist = &Artists
	if IsStartServer {
		LibArtists.SortingFilter = "name"
		LibArtists.Asc = true
		IsStartServer = false
	}
	template, errors := template.ParseFiles("static/html/libraryArtists.html")
	if errors != nil {
		fmt.Println("Error Parsing Template")
		fmt.Println(errors)
	}
	if r.Method == "GET" {
		setAllArtistVisibility(true)
	} else if r.Method == "POST" {
		searchContent := r.FormValue("searchBar")
		sortingOption := r.FormValue("sortFilter")
		sortingOrder := r.FormValue("sortOrder")
		if len(sortingOption) != 0 {
			LibArtists.SortingFilter = sortingOption
		}
		if len(searchContent) > 0 {
			searchArtists(searchContent)
		}
		if len(sortingOrder) != 0 {
			if sortingOrder == "asc" {
				LibArtists.Asc = true
			} else if sortingOrder == "desc" {
				LibArtists.Asc = false
			}
		}
	}
	sortArtists(LibArtists.SortingFilter, LibArtists.Asc)
	template.Execute(w, LibArtists)
}
