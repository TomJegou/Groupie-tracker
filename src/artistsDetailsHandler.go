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

func findArtistById(listArtist []Artist, id int) {
	for _, artist := range listArtist {
		if artist.Id == id {
			ChanArtDet <- artist
			return
		}
	}
	ChanArtDet <- listArtist[0]
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
		go findArtistById(Artists, idArtist)
		artist := <-ChanArtDet
		artistDetailled := ArtistDetailled{Artist: &artist, ArtistConcertsDatesLocation: Relations["index"][idArtist-1].DatesLocations}
		template.Execute(w, artistDetailled)
	}
}
