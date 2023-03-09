package handlers

import (
	"absolut-music/src/globalDataStructures"
	"absolut-music/src/structures"
	"absolut-music/src/tools"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

/*Artist detailled page's handler*/
func ArtistsDetailsHandler(w http.ResponseWriter, r *http.Request) {
	go tools.ChangeListenAddr(r)
	globalDataStructures.OnLibraryArtists = false
	// in case of restart server, check if the artist list is empty, if yes call the api
	if len(globalDataStructures.Artists) == 0 {
		var wg sync.WaitGroup
		wg.Add(1)
		go tools.PutBodyResponseApiIntoStruct(globalDataStructures.URLARTISTS, &globalDataStructures.Artists, &wg)
		wg.Wait()
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go tools.PutBodyResponseApiIntoStruct(globalDataStructures.URLRELATION, &globalDataStructures.Relations, &wg)
	wg.Wait()
	// get the artist's id and convert it into an int
	idArtist, err := strconv.Atoi(r.FormValue("artistCardId"))
	if err != nil {
		fmt.Println("Error converting string to integer")
		fmt.Println(err)
	} else {
		// Display the correct artist's information
		go tools.ParseHtml("static/html/artistsDetails.html")
		template := <-globalDataStructures.ChanTemplates
		go tools.FindArtistById(idArtist)
		artistDetailled := &structures.ArtistDetailled{Artist: <-globalDataStructures.ChanArtDet, ArtistConcertsDatesLocation: globalDataStructures.Relations["index"][idArtist-1].DatesLocations, ListenAddr: &globalDataStructures.ListeningAddr}
		template.Execute(w, artistDetailled)
	}
}
