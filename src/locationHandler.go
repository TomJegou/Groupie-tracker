package src

import "net/http"

func locationHandler(w http.ResponseWriter, r *http.Request) {
	ChangeListenAddr(r)
	OnLibraryArtists = false
	go ParseHtml("static/html/location.html")
	template := <-ChanTemplates
	template.Execute(w, ListeningAddr)
}
