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
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	locations    []string
	ConcertDates string
	Relations    string
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

type relation struct {
	Id            int
	DateLocations []string
}

func Image(id int) {
	var g []groupe
	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, err := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal([]byte(body), &g)
	if err != nil {
		fmt.Println("Error : 404")
		return
	}
	fmt.Println("", g[id].Image)
}

func GroupName() {
	var g []groupe
	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, err := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal([]byte(body), &g)
	if err != nil {
		fmt.Println("Error : 404")
		return
	}
	fmt.Println("Groupe :", g[3].Name)
}

func MembersName(id int) {
	var g []groupe
	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, err := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal([]byte(body), &g)
	if err != nil {
		fmt.Println("Error : 404")
		return
	}
	fmt.Println(g[id].Members)
}

func CreationDate(id int) {
	var g []groupe
	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, err := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal([]byte(body), &g)
	if err != nil {
		fmt.Println("Error : 404")
		return
	}
	fmt.Println(g[id].CreationDate)
}

func FirstAlbum(id int) {
	var g []groupe
	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, err := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal([]byte(body), &g)
	if err != nil {
		fmt.Println("Error : 404", err)
		return
	}
	fmt.Println(g[id].FirstAlbum)
}

func Locatlist(id int) {
	var l location
	url := "https://groupietrackers.herokuapp.com/api/locations/" + strconv.Itoa(id)
	req, err := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &l)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(l.Locations)
}

func Dates(id int) {
	var d date
	url := "https://groupietrackers.herokuapp.com/api/dates/" + strconv.Itoa(id)
	req, err := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &d)
	if err != nil {
		fmt.Println("Error 404")
		return
	}
	fmt.Println(d.Dates)
}

/* func Relations(id int) {
	var r []relation
	url := "https://groupietrackers.herokuapp.com/api/relation/" + strconv.Itoa(id)
	req, err := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println("Error 404", err)
		return
	}
	fmt.Println(r[id].dateLocations)
}
*/
func ArtistsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("static/html/artists.html")
	template.Execute(w, nil)
}
