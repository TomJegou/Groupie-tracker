package handlers

import (
	"net/http"
	"absolut-music/src/constances"
	"absolut-music/src/tools"
)

/*legal notice's handler*/
func LegalNoticeHandler(w http.ResponseWriter, r *http.Request) {
	go tools.ChangeListenAddr(r)
	constances.OnLibraryArtists = false
	go tools.ParseHtml("static/html/legalNotice.html")
	template := <-constances.ChanTemplates
	template.Execute(w, constances.ListeningAddr)
}
