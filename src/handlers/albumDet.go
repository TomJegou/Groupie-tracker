package handlers

import (
	"absolut-music/src/api"
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"net/http"
	"sync"
)

func AlbumDetHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.FormValue("albumCard")) > 0 {
		var wg sync.WaitGroup
		wg.Add(2)
		go tools.ChangeListenAddr(r, &wg)
		gds.OnLibraryArtists = false
		go tools.ParseHtml("static/html/albumDet.html")
		template := <-gds.ChanTemplates
		go api.PutBodyResponseApiIntoStruct(api.RequestApi(api.MakeReqAlbumDet(r.FormValue("albumCard"))), gds.AlbumDet.SpotifyAlbum, &wg)
		wg.Wait()
		template.Execute(w, gds.AlbumDet)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
