package tools

import (
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/structures"
	"fmt"
	"net/http"
	"strconv"
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
	gds.ChanTemplates <- template
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
func ChangeListenAddr(r *http.Request, wg *sync.WaitGroup) {
	defer wg.Done()
	if r.Host != gds.ListeningAddr.Ipv4 {
		gds.ListeningAddr.Ipv4 = r.Host
	}
}

/*
Find the artist who as the same id as the id passed as parameter
from the Artists slice
*/
func FindArtistById(id int) {
	for _, artist := range gds.Artists {
		if artist.Id == id {
			gds.ChanArtDet <- &artist
			return
		}
	}
	gds.ChanArtDet <- &gds.Artists[0]
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
	for i := 0; i < len(gds.Artists); i++ {
		SetArtistVisibility(&gds.Artists[i], isVisible)
	}
}

/*
Set all the artists visibility to false and search into the slice Artists
all the artists's name wich start with the same patern as the string
searchContent passed as parameter. Every artist found has his visibility set to true
*/
func SearchArtists(searchContent string) {
	SetAllArtistVisibility(false)
	for i := 0; i < len(gds.Artists); i++ {
		isOk := true
		for indexChar, char := range searchContent {
			if !strings.EqualFold(string(gds.Artists[i].Name[indexChar]), string(char)) {
				isOk = false
				break
			}
		}
		if isOk {
			SetArtistVisibility(&gds.Artists[i], true)
		}
	}
}

/*
Display all the artists from the slice Artists into pages
each page is put into the slice ListPages
*/
func DispatchIntoPage(wg *sync.WaitGroup) {
	defer wg.Done()
	gds.ListPages = []structures.Page{}
	pageCount := 0
	countArtist := 0
	page := structures.Page{Index: pageCount, Capacity: gds.PageCapacity, IsFirst: true}
	for i := 0; i < len(gds.Artists); i++ {
		if countArtist == gds.PageCapacity {
			gds.ListPages = append(gds.ListPages, page)
			pageCount++
			page = structures.Page{Index: pageCount, Capacity: gds.PageCapacity, IsFirst: false, IsLast: false}
			countArtist = 0
		}
		if gds.Artists[i].IsVisible {
			page.Content = append(page.Content, gds.Artists[i])
			countArtist++
		}
	}
	page.IsLast = true
	gds.ListPages = append(gds.ListPages, page)
}

/*Parse the first album's date into a structure FormatDate*/
func ParseDate(date string) structures.FormatDate {
	t := strings.Split(date, "-")
	yearInt, err := strconv.Atoi(t[2])
	if err != nil {
		fmt.Println(err)
	}
	monthInt, err := strconv.Atoi(t[1])
	if err != nil {
		fmt.Println(err)
	}
	dayInt, err := strconv.Atoi(t[0])
	if err != nil {
		fmt.Println(err)
	}
	parsedDate := structures.FormatDate{Year: yearInt, Month: monthInt, Day: dayInt}
	return parsedDate
}

/*Initialize the artists library*/
func InitLibArt() {
	if gds.IsStartServer {
		SetAllArtistVisibility(true)
		gds.LibArtists.SortingFilter = "name"
		gds.LibArtists.Asc = true
		QuickSort(gds.LibArtists.SortingFilter, gds.LibArtists.Asc)
		gds.LibArtists.IdPageToDisplay = 0
		gds.PageCapacity = 10
		RunParallel(DispatchIntoPage)
		gds.LibArtists.Page = &gds.ListPages[gds.LibArtists.IdPageToDisplay]
		gds.IsStartServer = false
	}
}

/*Creates a new LibLocations object and returns it's pointer*/
func NewLibLocations() *structures.LibLocations {
	locationList := make(map[string][]string)
	return &structures.LibLocations{LocationsList: locationList, ListenAddr: gds.ListeningAddr}
}

/*
Grabs in the Relations object all the cities and their concert dates in order to put them
into the libloca's attribute LocationsList wich is a map
*/
func GetLocations(libloca *structures.LibLocations, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(gds.Relations["index"]); i++ {
		for j := 0; j < len(gds.Relations["index"][i].DatesLocations); j++ {
			for cityName, listDate := range gds.Relations["index"][i].DatesLocations {
				if !libloca.InLocations(cityName) {
					libloca.LocationsList[cityName] = listDate
				} else {
					libloca.LocationsList[cityName] = append(libloca.LocationsList[cityName], listDate...)
				}
			}
		}
	}
}

/*Replace spaces to %20 for the spotify research querry*/
func PreprocessArtNameSearchSpotify(artistName string) string {
	result := ""
	l := strings.Split(artistName, " ")
	for index, kword := range l {
		if index != len(l)-1 {
			kword += "%20"
		}
		result += kword
	}
	return result
}

/*Check if all the artist from HeroKuapp are invisible*/
func CheckAllArtInvisible() bool {
	for _, art := range gds.Artists {
		if art.IsVisible {
			return false
		}
	}
	return true
}

func AppendtDate() {
	for _, dates := range gds.Dates["index"] {
		for _, date := range dates.Dates {
			if CheckDuplicateDate(date) {
				gds.DateHistr.Dates = append(gds.DateHistr.Dates, date)
			}
		}
	}
}

func CheckDuplicateDate(date string) bool {
	for _, t := range gds.DateHistr.Dates {
		if t == date {
			return false
		}
	}
	return true
}

func SortDates() {
	for i := 0; i < len(gds.DateHistr.Dates); i++ {
		x := i
		for z := i + 1; z < len(gds.DateHistr.Dates); z++ {
			DateX := ParseDate(gds.DateHistr.Dates[x])
			DateZ := ParseDate(gds.DateHistr.Dates[z])
			if DateZ.Year < DateX.Year {
				x = z
			} else if DateZ.Year == DateX.Year {
				if DateZ.Month < DateX.Month {
					x = z
				} else if DateZ.Month == DateX.Month {
					if DateZ.Day < DateX.Day {
						x = z
					}
				}
			}
		}
		gds.DateHistr.Dates[i], gds.DateHistr.Dates[x] = gds.DateHistr.Dates[x], gds.DateHistr.Dates[i]
	}
}
