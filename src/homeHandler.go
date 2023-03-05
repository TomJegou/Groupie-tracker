package src

import (
	"net/http"
)

/*Home page's handler*/
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	ListeningAddr.Ipv4 = r.Host
	OnLibraryArtists = false
	go ParseHtml("static/html/index.html")
	template := <-ChanTemplates
	template.Execute(w, ListeningAddr)
}
