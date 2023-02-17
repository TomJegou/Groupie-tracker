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

func findArtistById(listArtist []Artist, id int) (Artist, string) {
	for _, artist := range listArtist {
		if artist.Id == id {
			return artist, ""
		}
	}
	return listArtist[0], "Error id incorect"
}

func ArtistsDetailsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	idArtist, err := strconv.Atoi(r.FormValue("artistCard"))
	if err != nil {
		fmt.Println("Error converting string to integer")
		fmt.Println(err)
	}
	template, err := template.ParseFiles("static/html/artistsDetails.html")
	if err != nil {
		fmt.Println("Error parsing template artistsDetails.html")
		fmt.Println(err)
	}
	artist, errorId := findArtistById(Artists, idArtist)
	if errorId != "" {
		fmt.Println(errorId)
		http.Redirect(w, r, "/artists", http.StatusFound)
	} else {
		artistDetailled := ArtistDetailled{}
		artistDetailled.Artist = &artist
		artistDetailled.ArtistConcertsDatesLocation = Relations["index"][idArtist-1].DatesLocations // a changer
		template.Execute(w, artistDetailled)
	}

}
