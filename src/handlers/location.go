package handlers

import (
	"absolut-music/src/api"
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"net/http"
	"sync"
)

/*handles the Locations library*/
func LocationHandler(w http.ResponseWriter, r *http.Request) {
	/*Create new object Liblocation using the func NewLibLocations()*/
	var LibLocations = tools.NewLibLocations()
	gds.OnLibraryArtists = false
	var wg sync.WaitGroup
	wg.Add(2)
	go tools.ChangeListenAddr(r, &wg)
	go api.PutBodyResponseApiIntoStruct(api.RequestApi(api.MakeReqHerokuapp(gds.URLRELATION)), &gds.Relations, &wg)
	go tools.ParseHtml("static/html/locations.html")
	template := <-gds.ChanTemplates
	wg.Wait()
	wg.Add(1)
	go tools.GetLocations(LibLocations, &wg)
	wg.Wait()
	if len(r.FormValue("searchBar")) > 0 {
		tools.SearchBarLocate(r.FormValue("searchBar"), LibLocations)
	}
	template.Execute(w, LibLocations)
}
