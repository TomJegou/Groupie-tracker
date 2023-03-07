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
	for i := 0; i < len(lib.Locations()); i++ {
		if location == lib.Locations()[i] {
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
	ChangeListenAddr(r)
	libLocations.ListenAddr = &ListeningAddr
	OnLibraryArtists = false
	var wg sync.WaitGroup
	wg.Add(1)
	PutBodyResponseApiIntoStruct(URLRELATION, &Relations, &wg)
	wg.Wait()
	getLocations()
	go ParseHtml("static/html/locations.html")
	template := <-ChanTemplates
	template.Execute(w, libLocations)
}
