package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

type data struct {
	Artists   string
	Locations string
	Date      string
	Relation  string
}

type groupe struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	locations    string
	ConcertDates string
	Relations    string
}

func main() {
	art()
	loc()
	namelist()
	locatlist()
}

func art() {
	var d data
	url := "https://groupietrackers.herokuapp.com/api"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal([]byte(body), &d)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Println(d.Artists)
}

func loc() {
	var d data
	url := "https://groupietrackers.herokuapp.com/api"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal([]byte(body), &d)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Println(d.Locations)
}

func namelist() {
	var g []groupe
	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal([]byte(body), &g)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	for i := 0; i < len(g); i++ {
		fmt.Println("Groupe :", i, g[i].Name)
		fmt.Println(g[i].Members)
	}
}

func locatlist() {
	var g []groupe
	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal([]byte(body), &g)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	for i := 0; i < len(g); i++ {
		fmt.Println(g[i].locations)
	}
}

func ArtistsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("static/html/artists.html")
	template.Execute(w, nil)
}
