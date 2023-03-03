package src

import (
	"net/http"
)

type HistoricArtist struct {
	*Artist
	//ArtistConcertsDatesLocation map[string][]string
}

/*Historic page's handler*/
func HistoricHandler(w http.ResponseWriter, r *http.Request) {
	OnLibraryArtists = false
	go ParseHtml("static/html/historic.html")
	template := <-ChanTemplates
	template.Execute(w, nil)
}
