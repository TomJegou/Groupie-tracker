package handlers

import (
	"absolut-music/src/api"
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"net/http"
	"sync"
)

/*handles the Relation library*/
func ConcertDetailHandler(w http.ResponseWriter, r *http.Request) {
	gds.ConcertDetailPage.ListDate = nil
	gds.OnLibraryArtists = false
	var wg sync.WaitGroup
	wg.Add(2)
	// Change listen address and retrieve data from Heroku API in parallel
	go tools.ChangeListenAddr(r, &wg)
	go api.PutBodyResponseApiIntoStruct(api.RequestApi(api.MakeReqHerokuapp(gds.URLRELATION)), &gds.Relations, &wg)
	// Load HTML template
	go tools.ParseHtml("static/html/concertDet.html")
	template := <-gds.ChanTemplates
	// Wait for all goroutines to finish
	wg.Wait()
	// Extract concert dates for the specified city from the Relations data
	for _, relation := range gds.Relations["index"] {
		for city, date := range relation.DatesLocations {
			if city == r.FormValue("buttonCityName") {
				gds.ConcertDetailPage.ListDate = append(gds.ConcertDetailPage.ListDate, date...)
			}
		}
	}
	// Render the HTML template with the concert dates and other data
	template.Execute(w, gds.ConcertDetailPage)
}
