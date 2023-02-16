package src

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type ArtistDetailled struct {
	*Artist
	ArtistConcertsDatesLocation map[string][]string
}

func ArtistsDetailsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	idInt, err := strconv.Atoi(r.FormValue("artistCard"))
	if err != nil {
		fmt.Println("Error converting string to integer")
		fmt.Println(err)
	}
	template, err := template.ParseFiles("static/html/artistsDetails.html")
	if err != nil {
		fmt.Println("Error parsing template artistsDetails.html")
		fmt.Println(err)
	}
	artist := Artists[idInt]
	artistDetailled := ArtistDetailled{}
	artistDetailled.Artist = &artist
	artistDetailled.ArtistConcertsDatesLocation = Relations["index"][idInt-1].DatesLocations
	template.Execute(w, artistDetailled)
}
