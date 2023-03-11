package handlers

import (
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"net/http"
)

/*Home page's handler*/
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	go tools.ChangeListenAddr(r)
	gds.OnLibraryArtists = false
	go tools.ParseHtml("static/html/index.html")
	template := <-gds.ChanTemplates
	template.Execute(w, gds.ListeningAddr)
}
