package src

import (
	"encoding/json"
	"fmt"
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
	Locations    []string
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

func Accueil(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("static/html/index.html")
	template.Execute(w, nil)
}

var Artists []Artist
var Dates map[string][]Date
var Locations map[string][]Location
var Relations map[string][]Relation

func GetALlApi() {
	err := json.Unmarshal([]byte(GetApi(URLARTISTS)), &Artists)
	if err != nil {
		fmt.Println("Erreur Unmarshal JSON")
	}
	err = json.Unmarshal([]byte(GetApi(URLDATES)), &Dates)
	if err != nil {
		fmt.Println("Erreur Unmarshal JSON")
	}
	err = json.Unmarshal([]byte(GetApi(URLLOCATIONS)), &Locations)
	if err != nil {
		fmt.Println("Erreur Unmarshal JSON")
	}
	err = json.Unmarshal([]byte(GetApi(URLRELATION)), &Relations)
	if err != nil {
		fmt.Println("Erreur Unmarshal JSON")
	}
}

func StartServer() {
	GetALlApi()
	FileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", FileServer))
	http.HandleFunc("/", Accueil)
	http.HandleFunc("/artists", ArtistsHandlerFunc)
	http.HandleFunc("/artistsDetails", ArtistsDetailsHandlerFunc)
	http.HandleFunc("/about", AboutHandlerFunc)
	http.HandleFunc("/legalNotice", LegalNoticeHandlerFunc)
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
