package handlers

import (
	"absolut-music/src/api"
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"sync"
)

/*Handler func of the library artists*/
func LibraryArtistsHandler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	needSort := false
	needDispatch := false
	// call the api if the user wasn't in the libArt page
	if !gds.OnLibraryArtists {
		var wg sync.WaitGroup
		wg.Add(1)
		go api.PutBodyResponseApiIntoStruct(api.RequestApi(api.MakeReqHerokuapp(gds.URLARTISTS)), &gds.Artists, &wg)
		wg.Wait()
		gds.OnLibraryArtists = true
	}
	tools.InitLibArt()
	// If no filter or search criteria is applied, show all the artists
	if len(r.FormValue("searchBar")) == 0 && len(r.FormValue("sortFilter")) == 0 && len(r.FormValue("sortOrder")) == 0 && len(r.FormValue("pagination")) == 0 && len(r.FormValue("nbrElem")) == 0 {
		tools.SetAllArtistVisibility(true)
		needDispatch = true
		needSort = true
	} else {
		// change the number of elem to display
		if len(r.FormValue("nbrElem")) > 0 && r.FormValue("nbrElem") != gds.NumberOfElem {
			gds.NumberOfElem = r.FormValue("nbrElem")
			pageCapacityTmp, errors := strconv.Atoi(gds.NumberOfElem)
			gds.PageCapacity = pageCapacityTmp
			if errors != nil {
				fmt.Println(errors)
			}
			needDispatch = true
		}
		// change the page to display
		if len(r.FormValue("pagination")) > 0 {
			gds.PaginationRequest = r.FormValue("pagination")
			if len(gds.ListPages) > 0 {
				if gds.PaginationRequest == "next" {
					gds.LibArtists.IdPageToDisplay = int(math.Min(float64(len(gds.ListPages)-1), float64(gds.LibArtists.IdPageToDisplay+1)))
				} else {
					gds.LibArtists.IdPageToDisplay = int(math.Max(float64(0), float64(gds.LibArtists.IdPageToDisplay-1)))
				}
				gds.LibArtists.Page = &gds.ListPages[gds.LibArtists.IdPageToDisplay]
			} else {
				http.Redirect(w, r, "/libraryArtists", http.StatusFound)
			}
		}
		// change the artists sort filter
		if len(r.FormValue("sortFilter")) > 0 && r.FormValue("sortFilter") != gds.SortingOption {
			gds.SortingOption = r.FormValue("sortFilter")
			if gds.SortingOption == "NumberOfConcert" && len(gds.Dates["index"]) == 0 {
				var wg sync.WaitGroup
				wg.Add(1)
				go api.PutBodyResponseApiIntoStruct(api.RequestApi(api.MakeReqHerokuapp(gds.URLDATES)), &gds.Dates, &wg)
				wg.Wait()
			}
			gds.LibArtists.SortingFilter = gds.SortingOption
			needSort = true
		}
		// search the same artist's name patern as the string passed in the searching bar
		if len(r.FormValue("searchBar")) > 0 && r.FormValue("searchBar") != gds.SearchContent {
			gds.SearchContent = r.FormValue("searchBar")
			tools.SearchArtists(gds.SearchContent)
			if tools.CheckAllArtInvisible() {
				http.Redirect(w, r, "/spotify-search", http.StatusFound)
				return
			}
			needSort = true
			needDispatch = true
		}
		// change the artists order
		if len(r.FormValue("sortOrder")) > 0 && r.FormValue("sortOrder") != gds.SortingOrder {
			gds.SortingOrder = r.FormValue("sortOrder")
			if gds.SortingOrder == "asc" {
				gds.LibArtists.Asc = true
			} else if gds.SortingOrder == "desc" {
				gds.LibArtists.Asc = false
			}
			needSort = true
		}
	}
	wg.Add(1)
	go tools.ChangeListenAddr(r, &wg)
	go tools.ParseHtml("static/html/libraryArtists.html")
	template := <-gds.ChanTemplates
	// sort the artists list
	if needSort {
		if gds.LibArtists.SortingFilter == "NumberOfConcert" {
			tools.NumberOfConcert()
		} else {
			tools.QuickSort(gds.LibArtists.SortingFilter, gds.LibArtists.Asc)
		}
		needDispatch = true
	}
	// dispatch all the artists into pages
	if needDispatch {
		tools.RunParallel(tools.DispatchIntoPage)
		if gds.LibArtists.IdPageToDisplay > len(gds.ListPages)-1 {
			gds.LibArtists.IdPageToDisplay = len(gds.ListPages) - 1
		}
		gds.LibArtists.Page = &gds.ListPages[gds.LibArtists.IdPageToDisplay]
	}
	wg.Wait()
	template.Execute(w, gds.LibArtists)
}
