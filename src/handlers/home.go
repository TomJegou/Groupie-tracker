package handlers

import (
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"net/http"
	"sync"
)

/*Home page's handler*/
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	wg.Add(1)
	go tools.ChangeListenAddr(r, &wg)
	gds.OnLibraryArtists = false
	go tools.ParseHtml("static/html/index.html")
	template := <-gds.ChanTemplates
	wg.Wait()
	template.Execute(w, gds.ListeningAddr)
}
