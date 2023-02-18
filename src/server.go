package src

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const URLARTISTS = "https://groupietrackers.herokuapp.com/api/artists"
const URLDATES = "https://groupietrackers.herokuapp.com/api/dates"
const URLLOCATIONS = "https://groupietrackers.herokuapp.com/api/locations"
const URLRELATION = "https://groupietrackers.herokuapp.com/api/relation"

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

var Artists []Artist
var Dates map[string][]Date
var Locations map[string][]Location
var Relations map[string][]Relation

func GetApi(url string) string {
	req, errors := http.NewRequest("GET", url, nil)
	if errors != nil {
		fmt.Println("Error Request")
		fmt.Println(errors)
	}
	res, errors := http.DefaultClient.Do(req)
	if errors != nil {
		fmt.Println("Error default client")
		fmt.Println(errors)
	}
	defer res.Body.Close()
	body, errors := io.ReadAll(res.Body)
	if errors != nil {
		fmt.Println("Error during read body")
		fmt.Println(errors)
	}
	return string(body)
}

func PutBodyResponseApiIntoStruct(url string, structure interface{}) {
	err := json.Unmarshal([]byte(GetApi(url)), &structure)
	if err != nil {
		fmt.Println("Erreur Unmarshal JSON\n", err)
	}
}

func StartServer() {
	FileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", FileServer))
	http.HandleFunc("/", Accueil)
	http.HandleFunc("/artists", libraryArtists)
	http.HandleFunc("/artistsDetails", ArtistsDetailsHandlerFunc)
	http.HandleFunc("/about", AboutHandlerFunc)
	http.HandleFunc("/legalNotice", LegalNoticeHandlerFunc)
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
