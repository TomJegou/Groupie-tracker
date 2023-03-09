package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"absolut-music/src/structures"
	"absolut-music/src/constances"
	"absolut-music/src/tools"
)

/*Structures*/

type Page struct {
	Index    int
	IsFirst  bool
	IsLast   bool
	Capacity int
	Content  []structures.Artist
}

type LibraryArtists struct {
	Artistlist    *[]structures.Artist
	SortingFilter string
	Asc           bool
	*Page
	IdPageToDisplay int
	*structures.ListenAddr
}

/*Global variables*/

var PageCapacity int
var LibArtists LibraryArtists
var ListPages []Page

/*Functions*/

/*
Set the artists's attribute IsVisible to the isVisible boolean
passed as parameter
*/
func setArtistVisibility(a *structures.Artist, isVisible bool) {
	a.IsVisible = isVisible
}

/*
Set all of the artists's attribute IsVisible from the slice Artists
to the boolean isVisible passed as parameter
*/
func setAllArtistVisibility(isVisible bool) {
	for i := 0; i < len(constances.Artists); i++ {
		setArtistVisibility(&constances.Artists[i], isVisible)
	}
}

/*
Set all the artists visibility to false and search into the slice Artists
all the artists's name wich start with the same patern as the string
searchContent passed as parameter. Every artist found has his visibility set to true
*/
func searchArtists(searchContent string) {
	setAllArtistVisibility(false)
	for i := 0; i < len(constances.Artists); i++ {
		isOk := true
		for indexChar, char := range searchContent {
			if !strings.EqualFold(string(constances.Artists[i].Name[indexChar]), string(char)) {
				isOk = false
				break
			}
		}
		if isOk {
			setArtistVisibility(&constances.Artists[i], true)
		}
	}
}

/*
Display all the artists from the slice Artists into pages
each page is put into the slice ListPages
*/
func dispatchIntoPage(wg *sync.WaitGroup) {
	defer wg.Done()
	ListPages = []Page{}
	pageCount := 0
	countArtist := 0
	page := Page{Index: pageCount, Capacity: PageCapacity, IsFirst: true}
	for i := 0; i < len(constances.Artists); i++ {
		if countArtist == PageCapacity {
			ListPages = append(ListPages, page)
			pageCount++
			page = Page{Index: pageCount, Capacity: PageCapacity, IsFirst: false, IsLast: false}
			countArtist = 0
		}
		if constances.Artists[i].IsVisible {
			page.Content = append(page.Content, constances.Artists[i])
			countArtist++
		}
	}
	page.IsLast = true
	ListPages = append(ListPages, page)
}

func initLib() {
	if constances.IsStartServer {
		LibArtists.ListenAddr = &constances.ListeningAddr
		LibArtists.Artistlist = &constances.Artists
		setAllArtistVisibility(true)
		LibArtists.SortingFilter = "name"
		LibArtists.Asc = true
		tools.QuickSort(LibArtists.SortingFilter, LibArtists.Asc)
		LibArtists.IdPageToDisplay = 0
		PageCapacity = 10
		tools.RunParallel(dispatchIntoPage)
		LibArtists.Page = &ListPages[LibArtists.IdPageToDisplay]
		constances.IsStartServer = false
	}
}

/*Handler func of the library artists*/
func LibraryArtistsHandler(w http.ResponseWriter, r *http.Request) {
	go tools.ChangeListenAddr(r)
	needSort := false
	needDispatch := false
	if !constances.OnLibraryArtists {
		var wg sync.WaitGroup
		wg.Add(1)
		go tools.PutBodyResponseApiIntoStruct(constances.URLARTISTS, &constances.Artists, &wg)
		wg.Wait()
		constances.OnLibraryArtists = true
	}
	initLib()
	go tools.ParseHtml("static/html/libraryArtists.html")
	template := <-constances.ChanTemplates
	if r.Method == "GET" {
		setAllArtistVisibility(true)
		needDispatch = true
		needSort = true
	} else if r.Method == "POST" {
		searchContent := r.FormValue("searchBar")
		sortingOption := r.FormValue("sortFilter")
		sortingOrder := r.FormValue("sortOrder")
		paginationRequest := r.FormValue("pagination")
		numberOfElem := r.FormValue("nbrElem")
		if len(numberOfElem) != 0 {
			pageCapacityTmp, errors := strconv.Atoi(numberOfElem)
			PageCapacity = pageCapacityTmp
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
				LibArtists.Page = &ListPages[LibArtists.IdPageToDisplay]
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
		tools.QuickSort(LibArtists.SortingFilter, LibArtists.Asc)
		needDispatch = true
	}
	if needDispatch {
		tools.RunParallel(dispatchIntoPage)
		if LibArtists.IdPageToDisplay > len(ListPages)-1 {
			LibArtists.IdPageToDisplay = len(ListPages) - 1
		}
		LibArtists.Page = &ListPages[LibArtists.IdPageToDisplay]
	}
	template.Execute(w, LibArtists)
}
