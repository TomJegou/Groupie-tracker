package handlers

import (
	"absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"net/http"
)

/*About page's handler*/
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	go tools.ChangeListenAddr(r)
	globalDataStructures.OnLibraryArtists = false
	go tools.ParseHtml("static/html/about.html")
	template := <-globalDataStructures.ChanTemplates
	template.Execute(w, globalDataStructures.ListeningAddr)
}
