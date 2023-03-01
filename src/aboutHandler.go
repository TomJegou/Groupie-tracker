package src

import (
	"net/http"
)

func AboutHandlerFunc(w http.ResponseWriter, r *http.Request) {
	OnLibraryArtists = false
	go ParseHtml("static/html/about.html")
	template := <-ChanTemplates
	template.Execute(w, nil)
}
