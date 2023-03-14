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
	go tools.ChangeListenAddr(r, &wg)
	go api.PutBodyResponseApiIntoStruct(api.RequestApi(api.MakeReqHerokuapp(gds.URLRELATION)), &gds.Relations, &wg)
	go tools.ParseHtml("static/html/concertDet.html")
	template := <-gds.ChanTemplates
	wg.Wait()
	for _, relation := range gds.Relations["index"] {
		for city, date := range relation.DatesLocations {
			if city == r.FormValue("buttonCityName") {
				gds.ConcertDetailPage.ListDate = append(gds.ConcertDetailPage.ListDate, date...)
			}
		}
	}
	template.Execute(w, gds.ConcertDetailPage)
}
