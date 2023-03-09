package structures

/*Struct used for the locations library*/
type LibLocations struct {
	LocationsList map[string][]string
	*ListenAddr
}

/*Returns an array of all locations*/
func (lib LibLocations) Locations() []string {
	result := []string{}
	for cityName := range lib.LocationsList {
		result = append(result, cityName)
	}
	return result
}

/*Check if the location passed as a parameter is in the list of lactions
returns true if yes, else it returns false*/
func (lib LibLocations) InLocations(location string) bool {
	allLocation := lib.Locations()
	for i := 0; i < len(allLocation); i++ {
		if location == allLocation[i] {
			return true
		}
	}
	return false
}
