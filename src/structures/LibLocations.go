package structures

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
