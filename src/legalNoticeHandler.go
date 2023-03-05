package src

import (
	"net/http"
)

/*legal notice's handler*/
func LegalNoticeHandler(w http.ResponseWriter, r *http.Request) {
	ListeningAddr.Ipv4 = r.Host
	OnLibraryArtists = false
	go ParseHtml("static/html/legalNotice.html")
	template := <-ChanTemplates
	template.Execute(w, ListeningAddr)
}
