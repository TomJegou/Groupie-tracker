package src

import (
	"net/http"
)

/*legal notice's handler*/
func legalNoticeHandler(w http.ResponseWriter, r *http.Request) {
	ChangeListenAddr(r)
	OnLibraryArtists = false
	go ParseHtml("static/html/legalNotice.html")
	template := <-ChanTemplates
	template.Execute(w, ListeningAddr)
}
