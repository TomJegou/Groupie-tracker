package structures

/*Structures*/

type Artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
	IsVisible    bool
}

type Location struct {
	Id        int
	Locations []string
	Dates     string
}

type Date struct {
	Id    int
	Dates []string
}

type Relation struct {
	Id             int
	DatesLocations map[string][]string
}

type ListenAddr struct {
	Ipv4 string
	Port string
}
