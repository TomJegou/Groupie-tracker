package src

import (
	"net/http"
)

/*About page's handler*/
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	go ChangeListenAddr(r)
	OnLibraryArtists = false
	go ParseHtml("static/html/about.html")
	template := <-ChanTemplates
	template.Execute(w, ListeningAddr)
}
