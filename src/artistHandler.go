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

func main() {
	groupName(2)
	membersName(2)
	locatlist(2)
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

func ArtistsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("static/html/artists.html")
	template.Execute(w, nil)
}
