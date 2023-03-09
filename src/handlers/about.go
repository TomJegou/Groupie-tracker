package handlers

import (
	"absolut-music/src/constances"
	"absolut-music/src/tools"
	"net/http"
)

/*About page's handler*/
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	go tools.ChangeListenAddr(r)
	constances.OnLibraryArtists = false
	go tools.ParseHtml("static/html/about.html")
	template := <-constances.ChanTemplates
	template.Execute(w, constances.ListeningAddr)
}
