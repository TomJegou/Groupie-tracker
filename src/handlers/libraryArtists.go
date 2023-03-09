package handlers

import (
	"absolut-music/src/constances"
	"absolut-music/src/tools"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"sync"
)

/*Functions*/

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
	tools.InitLibArt()
	go tools.ParseHtml("static/html/libraryArtists.html")
	template := <-constances.ChanTemplates
	if r.Method == "GET" {
		tools.SetAllArtistVisibility(true)
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
			constances.PageCapacity = pageCapacityTmp
			if errors != nil {
				fmt.Println(errors)
			}
			needDispatch = true
		}
		if len(paginationRequest) != 0 {
			if len(constances.ListPages) > 0 {
				if paginationRequest == "next" {
					constances.LibArtists.IdPageToDisplay = int(math.Min(float64(len(constances.ListPages)-1), float64(constances.LibArtists.IdPageToDisplay+1)))
				} else {
					constances.LibArtists.IdPageToDisplay = int(math.Max(float64(0), float64(constances.LibArtists.IdPageToDisplay-1)))
				}
				constances.LibArtists.Page = &constances.ListPages[constances.LibArtists.IdPageToDisplay]
			} else {
				http.Redirect(w, r, "/libraryArtists", http.StatusFound)
			}
		}
		if len(sortingOption) != 0 {
			constances.LibArtists.SortingFilter = sortingOption
			needSort = true
		}
		if len(searchContent) > 0 {
			tools.SearchArtists(searchContent)
			needSort = true
			needDispatch = true
		}
		if len(sortingOrder) != 0 {
			if sortingOrder == "asc" {
				constances.LibArtists.Asc = true
			} else if sortingOrder == "desc" {
				constances.LibArtists.Asc = false
			}
			needSort = true
		}
	}
	if needSort {
		tools.QuickSort(constances.LibArtists.SortingFilter, constances.LibArtists.Asc)
		needDispatch = true
	}
	if needDispatch {
		tools.RunParallel(tools.DispatchIntoPage)
		if constances.LibArtists.IdPageToDisplay > len(constances.ListPages)-1 {
			constances.LibArtists.IdPageToDisplay = len(constances.ListPages) - 1
		}
		constances.LibArtists.Page = &constances.ListPages[constances.LibArtists.IdPageToDisplay]
	}
	template.Execute(w, constances.LibArtists)
}
