package src

import (
	"net/http"
)

/*Home page's handler*/
func homeHandler(w http.ResponseWriter, r *http.Request) {
	go ChangeListenAddr(r)
	OnLibraryArtists = false
	go ParseHtml("static/html/index.html")
	template := <-ChanTemplates
	template.Execute(w, ListeningAddr)
}
