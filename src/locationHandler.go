package src

import (
	"net/http"
	"sync"
)

type LibLocations struct {
	LocationsList map[string][]string
	*ListenAddr
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
	for i := 0; i < len(Relations["index"]); i++ {
		for j := 0; j < len(Relations["index"][i].DatesLocations); j++ {
			for cityName, listDate := range Relations["index"][i].DatesLocations {
				if !libLocations.InLocations(cityName) {
					libLocations.LocationsList[cityName] = listDate
				} else {
					libLocations.LocationsList[cityName] = append(libLocations.LocationsList[cityName], listDate...)
				}
			}
		}
	}
}

func locationHandler(w http.ResponseWriter, r *http.Request) {
	go ChangeListenAddr(r)
	libLocations.ListenAddr = &ListeningAddr
	OnLibraryArtists = false
	var wg sync.WaitGroup
	wg.Add(1)
	go PutBodyResponseApiIntoStruct(URLRELATION, &Relations, &wg)
	wg.Wait()
	go ParseHtml("static/html/locations.html")
	template := <-ChanTemplates
	getLocations()
	template.Execute(w, libLocations)
}
