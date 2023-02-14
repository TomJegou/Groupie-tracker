package src

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type ArtistDetailled struct {
	*Artist
	ArtistConcertsDates    []string
	ArtistConcertsLocation []string
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
	for _, location := range Locations["index"][idInt-1].Locations {
		artistDetailled.ArtistConcertsLocation = append(artistDetailled.ArtistConcertsLocation, location)
		artistDetailled.ArtistConcertsDates = append(artistDetailled.ArtistConcertsDates, Relations["index"][idInt-1].DatesLocations[location]...)
	}
	fmt.Println(artistDetailled.ArtistConcertsDates)
	template.Execute(w, artistDetailled)
}
