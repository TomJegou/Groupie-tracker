package handlers

import (
	"absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"sync"
)

/*Handler func of the library artists*/
func LibraryArtistsHandler(w http.ResponseWriter, r *http.Request) {
	go tools.ChangeListenAddr(r)
	needSort := false
	needDispatch := false
	// call the api if the user wasn't in the libArt page
	if !globalDataStructures.OnLibraryArtists {
		var wg sync.WaitGroup
		wg.Add(1)
		go tools.PutBodyResponseApiIntoStruct(globalDataStructures.URLARTISTS, &globalDataStructures.Artists, &wg)
		wg.Wait()
		globalDataStructures.OnLibraryArtists = true
	}
	tools.InitLibArt()
	go tools.ParseHtml("static/html/libraryArtists.html")
	template := <-globalDataStructures.ChanTemplates
	if r.Method == "GET" {
		//if the request's method is GET, set all the artist visible
		tools.SetAllArtistVisibility(true)
		needDispatch = true
		needSort = true
	} else if r.Method == "POST" {
		// if the request's method is POST, get all the values from the forms
		searchContent := r.FormValue("searchBar")
		sortingOption := r.FormValue("sortFilter")
		sortingOrder := r.FormValue("sortOrder")
		paginationRequest := r.FormValue("pagination")
		numberOfElem := r.FormValue("nbrElem")
		// change the number of elem to display
		if len(numberOfElem) != 0 {
			pageCapacityTmp, errors := strconv.Atoi(numberOfElem)
			globalDataStructures.PageCapacity = pageCapacityTmp
			if errors != nil {
				fmt.Println(errors)
			}
			needDispatch = true
		}
		// change the page to display
		if len(paginationRequest) != 0 {
			if len(globalDataStructures.ListPages) > 0 {
				if paginationRequest == "next" {
					globalDataStructures.LibArtists.IdPageToDisplay = int(math.Min(float64(len(globalDataStructures.ListPages)-1), float64(globalDataStructures.LibArtists.IdPageToDisplay+1)))
				} else {
					globalDataStructures.LibArtists.IdPageToDisplay = int(math.Max(float64(0), float64(globalDataStructures.LibArtists.IdPageToDisplay-1)))
				}
				globalDataStructures.LibArtists.Page = &globalDataStructures.ListPages[globalDataStructures.LibArtists.IdPageToDisplay]
			} else {
				http.Redirect(w, r, "/libraryArtists", http.StatusFound)
			}
		}
		// change the artists sort
		if len(sortingOption) != 0 {
			globalDataStructures.LibArtists.SortingFilter = sortingOption
			needSort = true
		}
		// search the same artist's name patern as the string passed in the searching bar
		if len(searchContent) > 0 {
			tools.SearchArtists(searchContent)
			needSort = true
			needDispatch = true
		}
		// change the artists order
		if len(sortingOrder) != 0 {
			if sortingOrder == "asc" {
				globalDataStructures.LibArtists.Asc = true
			} else if sortingOrder == "desc" {
				globalDataStructures.LibArtists.Asc = false
			}
			needSort = true
		}
	}
	// sort the artists list
	if needSort {
		tools.QuickSort(globalDataStructures.LibArtists.SortingFilter, globalDataStructures.LibArtists.Asc)
		needDispatch = true
	}
	// dispatch all the artists into pages
	if needDispatch {
		tools.RunParallel(tools.DispatchIntoPage)
		if globalDataStructures.LibArtists.IdPageToDisplay > len(globalDataStructures.ListPages)-1 {
			globalDataStructures.LibArtists.IdPageToDisplay = len(globalDataStructures.ListPages) - 1
		}
		globalDataStructures.LibArtists.Page = &globalDataStructures.ListPages[globalDataStructures.LibArtists.IdPageToDisplay]
	}
	template.Execute(w, globalDataStructures.LibArtists)
}
