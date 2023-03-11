package handlers

import (
	"absolut-music/src/api"
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"net/http"
	"sync"
)

func SpotifyHandler(w http.ResponseWriter, r *http.Request) {
	gds.SpotifyHdlStrct.ListenAddr = &gds.ListeningAddr
	var wg sync.WaitGroup
	wg.Add(2)
	go tools.ChangeListenAddr(r, &wg)
	gds.OnLibraryArtists = false
	go tools.ParseHtml("static/html/spotify.html")
	template := <-gds.ChanTemplates
	go api.PutBodyResponseApiIntoStruct(api.RequestApi(api.MakeReqSearchArtAPISportify(gds.SearchContent)), gds.ResultSpotifySearchArtist, &wg)
	wg.Wait()
	gds.SpotifyHdlStrct.SpotifySearchArtist = gds.ResultSpotifySearchArtist
	template.Execute(w, gds.SpotifyHdlStrct)
}
