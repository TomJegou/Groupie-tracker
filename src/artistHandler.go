package src

import (
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type Artists struct {
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

func getApi(url string) string {
	req, _ := http.NewRequest("GET", url, nil)
	res, errors := http.DefaultClient.Do(req)
	if errors != nil {
		log.Fatal(errors)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

func ArtistsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("static/html/artists.html")
	template.Execute(w, nil)
}
