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
	id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	locations    []string
	ConcertDates string
	Relations    string
	index        []string
}

type loca struct {
	Index     []string
	Id        int
	Locations []string
}

func main() {
	art()
	loc()
	groupName(2)
	membersName(2)
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

func groupName(id int) {
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
	fmt.Println("Groupe :", g[id].Name)
	fmt.Println(g[id].Members)
}

func membersName(id int) {
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
	fmt.Println(g[id].Members)
}

func locatlist() {
//louis
}

func ArtistsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("static/html/artists.html")
	template.Execute(w, nil)
}
