package handlers

import (
	"absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"net/http"
)

/*Home page's handler*/
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	go tools.ChangeListenAddr(r)
	globalDataStructures.OnLibraryArtists = false
	go tools.ParseHtml("static/html/index.html")
	template := <-globalDataStructures.ChanTemplates
	template.Execute(w, globalDataStructures.ListeningAddr)
}
