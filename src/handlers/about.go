package handlers

import (
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"net/http"
	"sync"
)

/*About page's handler*/
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	wg.Add(1)
	go tools.ChangeListenAddr(r, &wg)
	gds.OnLibraryArtists = false
	go tools.ParseHtml("static/html/about.html")
	template := <-gds.ChanTemplates
	wg.Wait()
	template.Execute(w, gds.ListeningAddr)
}
