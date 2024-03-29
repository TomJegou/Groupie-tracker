package handlers

import (
	"absolut-music/src/api"
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"net/http"
	"sync"
)

/*Hanler of the historic of the dates*/
func DateHistoricHandler(w http.ResponseWriter, r *http.Request) {
	gds.DateHistr.Dates = nil
	var wg sync.WaitGroup
	wg.Add(2)
	go tools.ChangeListenAddr(r, &wg)
	gds.OnLibraryArtists = false
	go tools.ParseHtml("static/html/dateHistoric.html")
	template := <-gds.ChanTemplates
	go api.PutBodyResponseApiIntoStruct(api.RequestApi(api.MakeReqHerokuapp(gds.URLDATES)), &gds.Dates, &wg)
	wg.Wait()
	tools.AppendtDate()
	tools.SortDates()
	template.Execute(w, gds.DateHistr)
}
