package src

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"text/template"
)

/*Structures*/

type Artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
	IsVisible    bool
}

type Location struct {
	Id        int
	Locations []string
}

type Date struct {
	Id    int
	Dates []string
}

type Relation struct {
	Id             int
	DatesLocations map[string][]string
}

/*Constances*/

const URLARTISTS = "https://groupietrackers.herokuapp.com/api/artists"
const URLDATES = "https://groupietrackers.herokuapp.com/api/dates"
const URLLOCATIONS = "https://groupietrackers.herokuapp.com/api/locations"
const URLRELATION = "https://groupietrackers.herokuapp.com/api/relation"

/*Variables*/

var IsStartServer = true
var OnLibraryArtists = false
var Artists []Artist
var Dates map[string][]Date
var Locations map[string][]Location
var Relations map[string][]Relation

/*Channels*/

var ChanArtists = make(chan *[]Artist)
var ChanTemplates = make(chan *template.Template)
var ChanArtDet = make(chan Artist)

/*Functions*/

/*Do an API call and return a string of the response*/
func GetApi(url string) string {
	req, errors := http.NewRequest("GET", url, nil)
	if errors != nil {
		fmt.Println("Error Request")
		fmt.Println(errors)
		return ""
	}
	res, errors := http.DefaultClient.Do(req)
	if errors != nil {
		fmt.Println("Error default client")
		fmt.Println(errors)
		return ""
	}
	defer res.Body.Close()
	body, errors := io.ReadAll(res.Body)
	if errors != nil {
		fmt.Println("Error during read body")
		fmt.Println(errors)
		return ""
	}
	return string(body)
}

/*
Call the API using the url passed as a parameter
and the func GetApi, and put the response into the structure passed as a parameter
*/
func PutBodyResponseApiIntoStruct(url string, structure interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	err := json.Unmarshal([]byte(GetApi(url)), &structure)
	if err != nil {
		fmt.Println("Erreur Unmarshal JSON\n", err)
	}
}

/*
Establish the routing for the webApp and start the server
on port 80
*/
func StartServer(wg *sync.WaitGroup) {
	defer wg.Done()
	FileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", FileServer))
	http.HandleFunc("/", Home)
	http.HandleFunc("/libraryArtists", libraryArtists)
	http.HandleFunc("/artistsDetails", ArtistsDetailsHandlerFunc)
	http.HandleFunc("/about", AboutHandlerFunc)
	http.HandleFunc("/legalNotice", LegalNoticeHandlerFunc)
	fmt.Println("http://localhost:80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error starting the server")
	}
}
