package handlers

import (
	"absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"absolut-music/src/api"
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
		go api.PutBodyResponseApiIntoStruct(api.RequestApi(api.MakeReqHerokuapp(globalDataStructures.URLARTISTS)), &globalDataStructures.Artists, &wg)
		wg.Wait()
		globalDataStructures.OnLibraryArtists = true
	}
	tools.InitLibArt()
	go tools.ParseHtml("static/html/libraryArtists.html")
	template := <-globalDataStructures.ChanTemplates
	if len(r.FormValue("searchBar")) == 0 && len(r.FormValue("sortFilter")) == 0 && len(r.FormValue("sortOrder")) == 0 && len(r.FormValue("pagination")) == 0 && len(r.FormValue("nbrElem")) == 0 {
		tools.SetAllArtistVisibility(true)
		needDispatch = true
		needSort = true
	} else {
		// change the number of elem to display
		if len(r.FormValue("nbrElem")) > 0 && r.FormValue("nbrElem") != globalDataStructures.NumberOfElem {
			globalDataStructures.NumberOfElem = r.FormValue("nbrElem")
			pageCapacityTmp, errors := strconv.Atoi(globalDataStructures.NumberOfElem)
			globalDataStructures.PageCapacity = pageCapacityTmp
			if errors != nil {
				fmt.Println(errors)
			}
			needDispatch = true
		}
		// change the page to display
		if len(r.FormValue("pagination")) > 0 {
			globalDataStructures.PaginationRequest = r.FormValue("pagination")
			if len(globalDataStructures.ListPages) > 0 {
				if globalDataStructures.PaginationRequest == "next" {
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
		if len(r.FormValue("sortFilter")) > 0 && r.FormValue("sortFilter") != globalDataStructures.SortingOption {
			globalDataStructures.SortingOption = r.FormValue("sortFilter")
			globalDataStructures.LibArtists.SortingFilter = globalDataStructures.SortingOption
			needSort = true
		}
		// search the same artist's name patern as the string passed in the searching bar
		if len(r.FormValue("searchBar")) > 0 && r.FormValue("searchBar") != globalDataStructures.SearchContent {
			globalDataStructures.SearchContent = r.FormValue("searchBar")
			tools.SearchArtists(globalDataStructures.SearchContent)
			needSort = true
			needDispatch = true
		}
		// change the artists order
		if len(r.FormValue("sortOrder")) > 0 && r.FormValue("sortOrder") != globalDataStructures.SortingOrder {
			globalDataStructures.SortingOrder = r.FormValue("sortOrder")
			if globalDataStructures.SortingOrder == "asc" {
				globalDataStructures.LibArtists.Asc = true
			} else if globalDataStructures.SortingOrder == "desc" {
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
