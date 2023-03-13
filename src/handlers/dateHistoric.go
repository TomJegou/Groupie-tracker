package handlers

import (
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"absolut-music/src/api"
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
	go api.PutBodyResponseApiIntoStruct()
	wg.Wait()
	tools.SortDate()
	fmt.Println(gds.DateHistr.Dates)
	template.Execute(w, gds.DateHistr)
}
