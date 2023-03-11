package handlers

import (
	"net/http"
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
)

/*legal notice's handler*/
func LegalNoticeHandler(w http.ResponseWriter, r *http.Request) {
	go tools.ChangeListenAddr(r)
	gds.OnLibraryArtists = false
	go tools.ParseHtml("static/html/legalNotice.html")
	template := <-gds.ChanTemplates
	template.Execute(w, gds.ListeningAddr)
}
