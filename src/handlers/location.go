package handlers

import (
	"absolut-music/src/constances"
	"absolut-music/src/structures"
	"absolut-music/src/tools"
	"net/http"
	"sync"
)

type LibLocations struct {
	LocationsList map[string][]string
	*structures.ListenAddr
}

func (lib LibLocations) Locations() []string {
	result := []string{}
	for cityName := range lib.LocationsList {
		result = append(result, cityName)
	}
	return result
}

func (lib LibLocations) InLocations(location string) bool {
	allLocation := lib.Locations()
	for i := 0; i < len(allLocation); i++ {
		if location == allLocation[i] {
			return true
		}
	}
	return false
}

func NewLibLocations() *LibLocations {
	locationList := make(map[string][]string)
	return &LibLocations{LocationsList: locationList}
}

var libLocations = NewLibLocations()

func getLocations() {
	for i := 0; i < len(constances.Relations["index"]); i++ {
		for j := 0; j < len(constances.Relations["index"][i].DatesLocations); j++ {
			for cityName, listDate := range constances.Relations["index"][i].DatesLocations {
				if !libLocations.InLocations(cityName) {
					libLocations.LocationsList[cityName] = listDate
				} else {
					libLocations.LocationsList[cityName] = append(libLocations.LocationsList[cityName], listDate...)
				}
			}
		}
	}
}

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	go tools.ChangeListenAddr(r)
	libLocations.ListenAddr = &constances.ListeningAddr
	constances.OnLibraryArtists = false
	var wg sync.WaitGroup
	wg.Add(1)
	go tools.PutBodyResponseApiIntoStruct(constances.URLRELATION, &constances.Relations, &wg)
	wg.Wait()
	go tools.ParseHtml("static/html/locations.html")
	template := <-constances.ChanTemplates
	getLocations()
	template.Execute(w, libLocations)
}
