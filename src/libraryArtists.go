package src

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

/*Structures*/

type Page struct {
	Index    int
	IsFirst  bool
	IsLast   bool
	Capacity int
	Content  []Artist
}

type LibraryArtists struct {
	Artistlist    *[]Artist
	SortingFilter string
	Asc           bool
	*Page
	IdPageToDisplay int
	*ListenAddr
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
func setArtistVisibility(a *Artist, isVisible bool) {
	a.IsVisible = isVisible
}

/*
Set all of the artists's attribute IsVisible from the slice Artists
to the boolean isVisible passed as parameter
*/
func setAllArtistVisibility(isVisible bool) {
	for i := 0; i < len(Artists); i++ {
		setArtistVisibility(&Artists[i], isVisible)
	}
}

/*
Set all the artists visibility to false and search into the slice Artists
all the artists's name wich start with the same patern as the string
searchContent passed as parameter. Every artist found has his visibility set to true
*/
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

/*Reverse the Artists slice*/
func reverseSliceArtist(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(Artists)/2; i++ {
		Artists[i], Artists[len(Artists)-1-i] = Artists[len(Artists)-1-i], Artists[i]
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

func SortFirstAlbum() {
	for i := 0; i < len(Artists); i++ {
		x := i
		for z := i + 1; z < len(Artists); z++ {
			splitx := strings.Split(Artists[x].FirstAlbum, "-")
			splitz := strings.Split(Artists[z].FirstAlbum, "-")
			YearX, err := strconv.Atoi(splitx[2])
			YearZ, err := strconv.Atoi(splitz[2])
			MonthX, err := strconv.Atoi(splitx[1])
			MonthZ, err := strconv.Atoi(splitz[1])
			DayX, err := strconv.Atoi(splitx[0])
			DayZ, err := strconv.Atoi(splitz[0])
			if err != nil {
				fmt.Println(err)
			}
			if YearZ < YearX {
				x = z
			} else if YearZ == YearX {
				if MonthZ < MonthX {
					x = z
				} else if MonthZ == MonthX {
					if DayZ < DayX {
						x = z
					}
				}
			}
		}
		Artists[i], Artists[x] = Artists[x], Artists[i]
	}
}

/*Handler func of the library artists*/
func LibraryArtistsHandler(w http.ResponseWriter, r *http.Request) {
	ChangeListenAddr(r)
	needSort := false
	needDispatch := false
	if !OnLibraryArtists {
		var wg sync.WaitGroup
		wg.Add(1)
		go PutBodyResponseApiIntoStruct(URLARTISTS, &Artists, &wg)
		wg.Wait()
		OnLibraryArtists = true
	}
	if IsStartServer {
		LibArtists.ListenAddr = &ListeningAddr
		LibArtists.Artistlist = &Artists
		setAllArtistVisibility(true)
		LibArtists.SortingFilter = "name"
		LibArtists.Asc = true
		QuickSort(LibArtists.SortingFilter, LibArtists.Asc)
		LibArtists.IdPageToDisplay = 0
		PageCapacity = 10
		RunParallel(dispatchIntoPage)
		LibArtists.Page = &ListPages[LibArtists.IdPageToDisplay]
		IsStartServer = false
	}
	go ParseHtml("static/html/libraryArtists.html")
	template := <-ChanTemplates
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
		QuickSort(LibArtists.SortingFilter, LibArtists.Asc)
		needDispatch = true
	}
	if needDispatch {
		RunParallel(dispatchIntoPage)
		if LibArtists.IdPageToDisplay > len(ListPages)-1 {
			LibArtists.IdPageToDisplay = len(ListPages) - 1
		}
		LibArtists.Page = &ListPages[LibArtists.IdPageToDisplay]
	}
	template.Execute(w, LibArtists)
}
