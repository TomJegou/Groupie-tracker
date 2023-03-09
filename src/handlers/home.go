package handlers

import (
	"absolut-music/src/constances"
	"absolut-music/src/tools"
	"net/http"
)

/*Home page's handler*/
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	go tools.ChangeListenAddr(r)
	constances.OnLibraryArtists = false
	go tools.ParseHtml("static/html/index.html")
	template := <-constances.ChanTemplates
	template.Execute(w, constances.ListeningAddr)
}
