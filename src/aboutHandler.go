package src

import (
	"net/http"
)

/*About page's handler*/
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	ListeningAddr.Ipv4 = r.Host
	OnLibraryArtists = false
	go ParseHtml("static/html/about.html")
	template := <-ChanTemplates
	template.Execute(w, ListeningAddr)
}
