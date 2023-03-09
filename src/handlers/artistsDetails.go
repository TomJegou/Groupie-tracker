package handlers

import (
	"absolut-music/src/constances"
	"absolut-music/src/structures"
	"absolut-music/src/tools"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

/*Structures*/

type ArtistDetailled struct {
	*structures.Artist
	ArtistConcertsDatesLocation map[string][]string
	*structures.ListenAddr
}

/*Functions*/

/*
Find the artist who as the same id as the id passed as parameter
from the Artists slice
*/
func findArtistById(id int) {
	for _, artist := range constances.Artists {
		if artist.Id == id {
			constances.ChanArtDet <- &artist
			return
		}
	}
	constances.ChanArtDet <- &constances.Artists[0]
}

/*Artist detailled page's handler*/
func ArtistsDetailsHandler(w http.ResponseWriter, r *http.Request) {
	go tools.ChangeListenAddr(r)
	constances.OnLibraryArtists = false
	if len(constances.Artists) == 0 {
		var wg sync.WaitGroup
		wg.Add(1)
		go tools.PutBodyResponseApiIntoStruct(constances.URLARTISTS, &constances.Artists, &wg)
		wg.Wait()
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go tools.PutBodyResponseApiIntoStruct(constances.URLRELATION, &constances.Relations, &wg)
	wg.Wait()
	idArtist, err := strconv.Atoi(r.FormValue("artistCardId"))
	if err != nil {
		fmt.Println("Error converting string to integer")
		fmt.Println(err)
	} else {
		go tools.ParseHtml("static/html/artistsDetails.html")
		template := <-constances.ChanTemplates
		go findArtistById(idArtist)
		artistDetailled := &ArtistDetailled{Artist: <-constances.ChanArtDet, ArtistConcertsDatesLocation: constances.Relations["index"][idArtist-1].DatesLocations, ListenAddr: &constances.ListeningAddr}
		template.Execute(w, artistDetailled)
	}
}
