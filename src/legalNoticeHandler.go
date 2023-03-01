package src

import (
	"net/http"
)

func LegalNoticeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	OnLibraryArtists = false
	go ParseHtml("static/html/legalNotice.html")
	template := <-ChanTemplates
	template.Execute(w, nil)
}
