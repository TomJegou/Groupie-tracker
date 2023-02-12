package src

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
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
	req, _ := http.NewRequest("GET", url, nil)
	res, errors := http.DefaultClient.Do(req)
	if errors != nil {
		log.Fatal(errors)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return string(body)
}

func GetALlApi() {
	err := json.Unmarshal([]byte(GetApi(URLARTISTS)), &Artists)
	if err != nil {
		fmt.Println("Erreur Unmarshal JSON artists")
		fmt.Println(err)
	}
	err = json.Unmarshal([]byte(GetApi(URLDATES)), &Dates)
	if err != nil {
		fmt.Println("Erreur Unmarshal JSON dates")
	}
	err = json.Unmarshal([]byte(GetApi(URLLOCATIONS)), &Locations)
	if err != nil {
		fmt.Println("Erreur Unmarshal JSON locations")
	}
	err = json.Unmarshal([]byte(GetApi(URLRELATION)), &Relations)
	if err != nil {
		fmt.Println("Erreur Unmarshal JSON relations")
	}
}

func StartServer() {
	GetALlApi()
	tmpl := template.Must(template.ParseFiles("libraryArtists.tmpl"))
	FileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", FileServer))
	http.HandleFunc("/", Accueil)
	http.HandleFunc("/artists", ArtistsHandlerFunc)
	tmpl.Execute(w, Artists)
	http.HandleFunc("/artistsDetails", ArtistsDetailsHandlerFunc)
	http.HandleFunc("/about", AboutHandlerFunc)
	http.HandleFunc("/legalNotice", LegalNoticeHandlerFunc)
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
