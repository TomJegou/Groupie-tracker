package handlers

import (
	"absolut-music/src/api"
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/structures"
	"absolut-music/src/tools"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

/*Artist detailled page's handler*/
func ArtistsDetailsHandler(w http.ResponseWriter, r *http.Request) {
	gds.OnLibraryArtists = false
	// in case of restart server, check if the artist list is empty, if yes call the api
	if len(gds.Artists) == 0 {
		var wg sync.WaitGroup
		wg.Add(1)
		go api.PutBodyResponseApiIntoStruct(api.RequestApi(api.MakeReqHerokuapp(gds.URLARTISTS)), &gds.Artists, &wg)
		wg.Wait()
	}
	// get the artist's id and convert it into an int
	if len(r.FormValue("artistCardId")) > 0 {
		idArtist, err := strconv.Atoi(r.FormValue("artistCardId"))
		if err != nil {
			fmt.Println("Error converting string to integer")
			fmt.Println(err)
		} else {
			// Display the correct artist's information
			var wg sync.WaitGroup
			wg.Add(4)
			go tools.ChangeListenAddr(r, &wg)
			go api.PutBodyResponseApiIntoStruct(api.RequestApi(api.MakeReqHerokuapp(gds.URLRELATION)), &gds.Relations, &wg)
			go tools.ParseHtml("static/html/artistsDetails.html")
			template := <-gds.ChanTemplates
			go tools.FindArtistById(idArtist, &wg)
			artistDetailled := &structures.ArtistDetailled{Artist: <-gds.ChanArtDet, ArtistConcertsDatesLocation: gds.Relations["index"][idArtist-1].DatesLocations, ListenAddr: &gds.ListeningAddr}
			api.PutBodyResponseApiIntoStruct(api.RequestApi(api.MakeReqSearchArtAPISportify(artistDetailled.Name)), gds.ResultSpotifySearchArtist, &wg)
			//api.PutBodyResponseApiIntoStruct(api.RequestApi(api.MakeReqSearchAlbumArtAPISportify(artistDetailled.Name)), gds.ResultSpotifySearchAlbum, &wg)
			wg.Wait()
			artistDetailled.SpotifySearchArtist = gds.ResultSpotifySearchArtist
			//fmt.Println(gds.ResultSpotifySearchAlbum)
			template.Execute(w, artistDetailled)
		}
	} else {
		http.Redirect(w, r, "/libraryArtists", http.StatusFound)
	}
}
