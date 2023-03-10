package handlers

import (
	"net/http"
	"absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
)

/*legal notice's handler*/
func LegalNoticeHandler(w http.ResponseWriter, r *http.Request) {
	go tools.ChangeListenAddr(r)
	globalDataStructures.OnLibraryArtists = false
	go tools.ParseHtml("static/html/legalNotice.html")
	template := <-globalDataStructures.ChanTemplates
	template.Execute(w, globalDataStructures.ListeningAddr)
}
