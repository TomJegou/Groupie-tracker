package src

import (
	"fmt"
	"net/http"
	"strconv"
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
	OnLibraryArtists = false
	if len(Artists) == 0 {
		PutBodyResponseApiIntoStruct(URLARTISTS, &Artists)
	}
	PutBodyResponseApiIntoStruct(URLRELATION, &Relations)
	idArtist, err := strconv.Atoi(r.FormValue("artistCardId"))
	if err != nil {
		fmt.Println("Error converting string to integer")
		fmt.Println(err)
	} else {
		go ParseHtml("static/html/artistsDetails.html")
		template := <-ChanTemplates
		template.Execute(w, nil)
		artist, errorId := findArtistById(Artists, idArtist)
		if errorId != "" {
			fmt.Println(errorId)
			http.Redirect(w, r, "/libraryArtists", http.StatusFound)
		} else {
			artistDetailled := ArtistDetailled{Artist: &artist, ArtistConcertsDatesLocation: Relations["index"][idArtist-1].DatesLocations}
			template.Execute(w, artistDetailled)
		}
	}
}
