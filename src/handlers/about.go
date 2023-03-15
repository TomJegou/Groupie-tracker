package handlers

import (
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"net/http"
	"sync"
)

/*About page's handler*/
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	// Create a WaitGroup to synchronize the execution of goroutines
	var wg sync.WaitGroup
	wg.Add(1)
	// Launch a goroutine to update the server's listening address
	go tools.ChangeListenAddr(r, &wg)
	gds.OnLibraryArtists = false
	// Launch a goroutine to parse the about.html template
	go tools.ParseHtml("static/html/about.html")
	// Retrieve the parsed template from the channel
	template := <-gds.ChanTemplates
	// Wait for the goroutine to finish updating the listening address
	wg.Wait()
	// Execute the template and write the response
	template.Execute(w, gds.ListeningAddr)
}
