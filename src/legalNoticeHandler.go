package src

import (
	"net/http"
)

func LegalNoticeHandler(w http.ResponseWriter, r *http.Request) {
	OnLibraryArtists = false
	go ParseHtml("static/html/legalNotice.html")
	template := <-ChanTemplates
	template.Execute(w, nil)
}
