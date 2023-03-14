package handlers

import (
	"absolut-music/src/api"
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"net/http"
	"sync"
)

/*
Handler for the spotify artist when searchin for an artist who'isn't in
the herokuapp api
*/
func SpotifyHandler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	if len(r.FormValue("searchBarSpotify")) > 0 {
		gds.SearchContent = r.FormValue("searchBarSpotify")
	}
	wg.Add(2)
	go tools.ChangeListenAddr(r, &wg)
	gds.OnLibraryArtists = false
	go tools.ParseHtml("static/html/spotify.html")
	template := <-gds.ChanTemplates
	go api.PutBodyResponseApiIntoStruct(api.RequestApi(api.MakeReqSearchArtAPISportify(gds.SearchContent, "1")), gds.ResultSpotifySearchArtist, &wg)
	wg.Wait()
	gds.SpotifyHdlStrct.SpotifySearchArtist = gds.ResultSpotifySearchArtist
	gds.SearchContent = ""
	template.Execute(w, gds.SpotifyHdlStrct)
}
