package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

type location struct {
	Index     []string
	Id        int
	Locations []string
}

type date struct {
	Id    int
	Dates []string
}

func main() {
	image(1)
	groupName(1)
	membersName(1)
	FirstAlbum(1)
	locatlist(1)
	dates(1)
}

func image(id int) {
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
	fmt.Println("", g[id].Image)
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

func FirstAlbum(id int) {
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
	fmt.Println(g[id].FirstAlbum)
}

func locatlist(id int) {
	var l location
	url := "https://groupietrackers.herokuapp.com/api/locations/" + strconv.Itoa(id)
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(body, &l)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(l.Locations)
}

func dates(id int) {
	var d date
	url := "https://groupietrackers.herokuapp.com/api/dates/" + strconv.Itoa(id)
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(body, &d)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(d.Dates)
}

func ArtistsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("static/html/artists.html")
	template.Execute(w, nil)
}
