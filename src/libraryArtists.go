package src

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

type Page struct {
	Index    int
	IsFirst  bool
	IsLast   bool
	Capacity int
	Content  []Artist
}

type LibraryArtists struct {
	Artistlist      *[]Artist
	SortingFilter   string
	Asc             bool
	ThePage         *Page
	IdPageToDisplay int
}

var PageCapacity int
var LibArtists LibraryArtists
var ListPages []Page

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

func selectionSort() {

}

func dispatchIntoPage() {
	ListPages = []Page{}
	pageCount := 0
	countArtist := 0
	page := Page{Index: pageCount, Capacity: PageCapacity, IsFirst: true}
	for i := 0; i < len(Artists); i++ {
		if countArtist == PageCapacity {
			ListPages = append(ListPages, page)
			pageCount++
			page = Page{Index: pageCount, Capacity: PageCapacity, IsFirst: false, IsLast: false}
			countArtist = 0
		}
		if Artists[i].IsVisible {
			page.Content = append(page.Content, Artists[i])
			countArtist++
		}
	}
	page.IsLast = true
	ListPages = append(ListPages, page)
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
				//SortFirstAlbum()
			}
		}
		Artists[i], Artists[x] = Artists[x], Artists[i]
	}
	if !asc {
		reverseSliceArtist()
	}
}

func libraryArtists(w http.ResponseWriter, r *http.Request) {
	needSort := false
	needDispatch := false
	if !OnLibraryArtists {
		PutBodyResponseApiIntoStruct(URLARTISTS, &Artists)
		OnLibraryArtists = true
	}
	if IsStartServer {
		LibArtists.Artistlist = &Artists
		setAllArtistVisibility(true)
		LibArtists.SortingFilter = "name"
		LibArtists.Asc = true
		sortArtists(LibArtists.SortingFilter, LibArtists.Asc)
		LibArtists.IdPageToDisplay = 0
		PageCapacity = 10
		dispatchIntoPage()
		LibArtists.ThePage = &ListPages[LibArtists.IdPageToDisplay]
		IsStartServer = false
	}
	template, errors := template.ParseFiles("static/html/libraryArtists.html")
	if errors != nil {
		fmt.Println("Error Parsing Template")
		fmt.Println(errors)
	}
	if r.Method == "GET" {
		setAllArtistVisibility(true)
		needDispatch = true
	} else if r.Method == "POST" {
		searchContent := r.FormValue("searchBar")
		sortingOption := r.FormValue("sortFilter")
		sortingOrder := r.FormValue("sortOrder")
		paginationRequest := r.FormValue("pagination")
		numberOfElem := r.FormValue("nbrElem")
		if len(numberOfElem) != 0 {
			PageCapacity, errors = strconv.Atoi(numberOfElem)
			if errors != nil {
				fmt.Println(errors)
			}
			needDispatch = true
		}
		if len(paginationRequest) != 0 {
			if len(ListPages) > 0 {
				if paginationRequest == "next" {
					LibArtists.IdPageToDisplay = int(math.Min(float64(len(ListPages)-1), float64(LibArtists.IdPageToDisplay+1)))
				} else {
					LibArtists.IdPageToDisplay = int(math.Max(float64(0), float64(LibArtists.IdPageToDisplay-1)))
				}
				LibArtists.ThePage = &ListPages[LibArtists.IdPageToDisplay]
			} else {
				http.Redirect(w, r, "/libraryArtists", http.StatusFound)
			}
		}
		if len(sortingOption) != 0 {
			LibArtists.SortingFilter = sortingOption
			needSort = true
		}
		if len(searchContent) > 0 {
			searchArtists(searchContent)
			needSort = true
			needDispatch = true
		}
		if len(sortingOrder) != 0 {
			if sortingOrder == "asc" {
				LibArtists.Asc = true
			} else if sortingOrder == "desc" {
				LibArtists.Asc = false
			}
			needSort = true
		}
	}
	if needSort {
		sortArtists(LibArtists.SortingFilter, LibArtists.Asc)
		needDispatch = true
	}
	if needDispatch {
		dispatchIntoPage()
		LibArtists.ThePage = &ListPages[LibArtists.IdPageToDisplay]
	}
	template.Execute(w, LibArtists)
}
