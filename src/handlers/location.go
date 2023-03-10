package handlers

import (
	"absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"net/http"
	"sync"
)

/*Create new object Liblocation using the func NewLibLocations()*/
var LibLocations = tools.NewLibLocations()

/*handles the Locations library*/
func LocationHandler(w http.ResponseWriter, r *http.Request) {
	go tools.ChangeListenAddr(r)
	LibLocations.ListenAddr = &globalDataStructures.ListeningAddr
	globalDataStructures.OnLibraryArtists = false
	var wg sync.WaitGroup
	wg.Add(1)
	go tools.PutBodyResponseApiIntoStruct(tools.RequestApi(tools.MakeReqHerokuapp(globalDataStructures.URLRELATION)), &globalDataStructures.Relations, &wg)
	wg.Wait()
	go tools.ParseHtml("static/html/locations.html")
	template := <-globalDataStructures.ChanTemplates
	tools.GetLocations(LibLocations)
	template.Execute(w, LibLocations)
}
