package handlers

import (
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"net/http"
)

/*About page's handler*/
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	go tools.ChangeListenAddr(r)
	gds.OnLibraryArtists = false
	go tools.ParseHtml("static/html/about.html")
	template := <-gds.ChanTemplates
	template.Execute(w, gds.ListeningAddr)
}
