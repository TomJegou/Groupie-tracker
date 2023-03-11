package handlers

import (
	"absolut-music/src/api"
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"net/http"
	"sync"
)

/*Create new object Liblocation using the func NewLibLocations()*/
var LibLocations = tools.NewLibLocations()

/*handles the Locations library*/
func LocationHandler(w http.ResponseWriter, r *http.Request) {
	go tools.ChangeListenAddr(r)
	LibLocations.ListenAddr = &gds.ListeningAddr
	gds.OnLibraryArtists = false
	var wg sync.WaitGroup
	wg.Add(1)
	go api.PutBodyResponseApiIntoStruct(api.RequestApi(api.MakeReqHerokuapp(gds.URLRELATION)), &gds.Relations, &wg)
	wg.Wait()
	go tools.ParseHtml("static/html/locations.html")
	template := <-gds.ChanTemplates
	tools.GetLocations(LibLocations)
	template.Execute(w, LibLocations)
}
