package tools

import (
	"absolut-music/src/globalDataStructures"
	"absolut-music/src/structures"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"text/template"
)

/*
Parse the html file passed as a parameter and send the
template into  the ChanTemplates Channel
*/
func ParseHtml(fileToParse string) {
	template, errors := template.ParseFiles(fileToParse)
	if errors != nil {
		fmt.Println("Error Parsing Template")
		fmt.Println(errors)
	}
	globalDataStructures.ChanTemplates <- template
}

/*Make the function passed as a parameter run in Parallel as a goroutine*/
func RunParallel(f func(*sync.WaitGroup)) {
	var wg sync.WaitGroup
	wg.Add(1)
	go f(&wg)
	wg.Wait()
}

/*
Check if in the request if the host ipv4 is
the same as the one to be used for the templates.
If it's not the same, change the ListeningAddr.Ipv4 to the host requested
*/
func ChangeListenAddr(r *http.Request) {
	if r.Host != globalDataStructures.ListeningAddr.Ipv4 {
		globalDataStructures.ListeningAddr.Ipv4 = r.Host
	}
}

/*
Find the artist who as the same id as the id passed as parameter
from the Artists slice
*/
func FindArtistById(id int) {
	for _, artist := range globalDataStructures.Artists {
		if artist.Id == id {
			globalDataStructures.ChanArtDet <- &artist
			return
		}
	}
	globalDataStructures.ChanArtDet <- &globalDataStructures.Artists[0]
}

/*
Set the artists's attribute IsVisible to the isVisible boolean
passed as parameter
*/
func SetArtistVisibility(a *structures.Artist, isVisible bool) {
	a.IsVisible = isVisible
}

/*
Set all of the artists's attribute IsVisible from the slice Artists
to the boolean isVisible passed as parameter
*/
func SetAllArtistVisibility(isVisible bool) {
	for i := 0; i < len(globalDataStructures.Artists); i++ {
		SetArtistVisibility(&globalDataStructures.Artists[i], isVisible)
	}
}

/*
Set all the artists visibility to false and search into the slice Artists
all the artists's name wich start with the same patern as the string
searchContent passed as parameter. Every artist found has his visibility set to true
*/
func SearchArtists(searchContent string) {
	SetAllArtistVisibility(false)
	for i := 0; i < len(globalDataStructures.Artists); i++ {
		isOk := true
		for indexChar, char := range searchContent {
			if !strings.EqualFold(string(globalDataStructures.Artists[i].Name[indexChar]), string(char)) {
				isOk = false
				break
			}
		}
		if isOk {
			SetArtistVisibility(&globalDataStructures.Artists[i], true)
		}
	}
}

/*
Display all the artists from the slice Artists into pages
each page is put into the slice ListPages
*/
func DispatchIntoPage(wg *sync.WaitGroup) {
	defer wg.Done()
	globalDataStructures.ListPages = []structures.Page{}
	pageCount := 0
	countArtist := 0
	page := structures.Page{Index: pageCount, Capacity: globalDataStructures.PageCapacity, IsFirst: true}
	for i := 0; i < len(globalDataStructures.Artists); i++ {
		if countArtist == globalDataStructures.PageCapacity {
			globalDataStructures.ListPages = append(globalDataStructures.ListPages, page)
			pageCount++
			page = structures.Page{Index: pageCount, Capacity: globalDataStructures.PageCapacity, IsFirst: false, IsLast: false}
			countArtist = 0
		}
		if globalDataStructures.Artists[i].IsVisible {
			page.Content = append(page.Content, globalDataStructures.Artists[i])
			countArtist++
		}
	}
	page.IsLast = true
	globalDataStructures.ListPages = append(globalDataStructures.ListPages, page)
}

func InitLibArt() {
	if globalDataStructures.IsStartServer {
		globalDataStructures.LibArtists.ListenAddr = &globalDataStructures.ListeningAddr
		globalDataStructures.LibArtists.Artistlist = &globalDataStructures.Artists
		SetAllArtistVisibility(true)
		globalDataStructures.LibArtists.SortingFilter = "name"
		globalDataStructures.LibArtists.Asc = true
		QuickSort(globalDataStructures.LibArtists.SortingFilter, globalDataStructures.LibArtists.Asc)
		globalDataStructures.LibArtists.IdPageToDisplay = 0
		globalDataStructures.PageCapacity = 10
		RunParallel(DispatchIntoPage)
		globalDataStructures.LibArtists.Page = &globalDataStructures.ListPages[globalDataStructures.LibArtists.IdPageToDisplay]
		globalDataStructures.IsStartServer = false
	}
}

func NewLibLocations() *structures.LibLocations {
	locationList := make(map[string][]string)
	return &structures.LibLocations{LocationsList: locationList}
}

func GetLocations(libloca *structures.LibLocations) {
	for i := 0; i < len(globalDataStructures.Relations["index"]); i++ {
		for j := 0; j < len(globalDataStructures.Relations["index"][i].DatesLocations); j++ {
			for cityName, listDate := range globalDataStructures.Relations["index"][i].DatesLocations {
				if !libloca.InLocations(cityName) {
					libloca.LocationsList[cityName] = listDate
				} else {
					libloca.LocationsList[cityName] = append(libloca.LocationsList[cityName], listDate...)
				}
			}
		}
	}
}
