package src

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

/*Structures*/

type ArtistDetailled struct {
	*Artist
	ArtistConcertsDatesLocation map[string][]string
}

/*Functions*/

/*
Find the artist who as the same id as the id passed as parameter
from the Artists slice
*/
func findArtistById(id int) {
	for _, artist := range Artists {
		if artist.Id == id {
			ChanArtDet <- artist
			return
		}
	}
	ChanArtDet <- Artists[0]
}

/*Artist detailled page's handler*/
func ArtistsDetailsHandler(w http.ResponseWriter, r *http.Request) {
	OnLibraryArtists = false
	if len(Artists) == 0 {
		var wg sync.WaitGroup
		wg.Add(1)
		go PutBodyResponseApiIntoStruct(URLARTISTS, &Artists, &wg)
		wg.Wait()
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go PutBodyResponseApiIntoStruct(URLRELATION, &Relations, &wg)
	wg.Wait()
	idArtist, err := strconv.Atoi(r.FormValue("artistCardId"))
	if err != nil {
		fmt.Println("Error converting string to integer")
		fmt.Println(err)
	} else {
		go ParseHtml("static/html/artistsDetails.html")
		template := <-ChanTemplates
		go findArtistById(idArtist)
		artist := <-ChanArtDet
		artistDetailled := ArtistDetailled{Artist: &artist, ArtistConcertsDatesLocation: Relations["index"][idArtist-1].DatesLocations}
		template.Execute(w, artistDetailled)
	}
}
