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
	gds.OnLibraryArtists = false
	var wg sync.WaitGroup
	wg.Add(3)
	go tools.ChangeListenAddr(r, &wg)
	LibLocations.ListenAddr = &gds.ListeningAddr
	go api.PutBodyResponseApiIntoStruct(api.RequestApi(api.MakeReqHerokuapp(gds.URLRELATION)), &gds.Relations, &wg)
	go tools.ParseHtml("static/html/locations.html")
	template := <-gds.ChanTemplates
	tools.GetLocations(LibLocations, &wg)
	wg.Wait()
	template.Execute(w, LibLocations)
}
