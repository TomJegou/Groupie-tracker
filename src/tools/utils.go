package tools

import (
	"absolut-music/src/constances"
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
	constances.ChanTemplates <- template
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
	if r.Host != constances.ListeningAddr.Ipv4 {
		constances.ListeningAddr.Ipv4 = r.Host
	}
}

/*
Find the artist who as the same id as the id passed as parameter
from the Artists slice
*/
func FindArtistById(id int) {
	for _, artist := range constances.Artists {
		if artist.Id == id {
			constances.ChanArtDet <- &artist
			return
		}
	}
	constances.ChanArtDet <- &constances.Artists[0]
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
	for i := 0; i < len(constances.Artists); i++ {
		SetArtistVisibility(&constances.Artists[i], isVisible)
	}
}

/*
Set all the artists visibility to false and search into the slice Artists
all the artists's name wich start with the same patern as the string
searchContent passed as parameter. Every artist found has his visibility set to true
*/
func SearchArtists(searchContent string) {
	SetAllArtistVisibility(false)
	for i := 0; i < len(constances.Artists); i++ {
		isOk := true
		for indexChar, char := range searchContent {
			if !strings.EqualFold(string(constances.Artists[i].Name[indexChar]), string(char)) {
				isOk = false
				break
			}
		}
		if isOk {
			SetArtistVisibility(&constances.Artists[i], true)
		}
	}
}

/*
Display all the artists from the slice Artists into pages
each page is put into the slice ListPages
*/
func DispatchIntoPage(wg *sync.WaitGroup) {
	defer wg.Done()
	constances.ListPages = []structures.Page{}
	pageCount := 0
	countArtist := 0
	page := structures.Page{Index: pageCount, Capacity: constances.PageCapacity, IsFirst: true}
	for i := 0; i < len(constances.Artists); i++ {
		if countArtist == constances.PageCapacity {
			constances.ListPages = append(constances.ListPages, page)
			pageCount++
			page = structures.Page{Index: pageCount, Capacity: constances.PageCapacity, IsFirst: false, IsLast: false}
			countArtist = 0
		}
		if constances.Artists[i].IsVisible {
			page.Content = append(page.Content, constances.Artists[i])
			countArtist++
		}
	}
	page.IsLast = true
	constances.ListPages = append(constances.ListPages, page)
}

func InitLibArt() {
	if constances.IsStartServer {
		constances.LibArtists.ListenAddr = &constances.ListeningAddr
		constances.LibArtists.Artistlist = &constances.Artists
		SetAllArtistVisibility(true)
		constances.LibArtists.SortingFilter = "name"
		constances.LibArtists.Asc = true
		QuickSort(constances.LibArtists.SortingFilter, constances.LibArtists.Asc)
		constances.LibArtists.IdPageToDisplay = 0
		constances.PageCapacity = 10
		RunParallel(DispatchIntoPage)
		constances.LibArtists.Page = &constances.ListPages[constances.LibArtists.IdPageToDisplay]
		constances.IsStartServer = false
	}
}
