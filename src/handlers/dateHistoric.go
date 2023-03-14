package handlers

import (
	"absolut-music/src/api"
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"fmt"
	"net/http"
	"sync"
)

func DateHistoricHandler(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println(gds.DateHistr.Dates[0])
	template.Execute(w, gds.DateHistr)
}
