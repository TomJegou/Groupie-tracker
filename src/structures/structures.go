package structures

/*Struct used to store the data from the api*/
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

/*Struct used to store the listening address*/
type ListenAddr struct {
	Ipv4 string
	Port string
}

/*Struct used for the pageg Artists detailled*/
type ArtistDetailled struct {
	*Artist
	ArtistConcertsDatesLocation map[string][]string
	*ListenAddr
}

/*Structs used for the Artists library*/
type Page struct {
	Index    int
	IsFirst  bool
	IsLast   bool
	Capacity int
	Content  []Artist
}

type LibraryArtists struct {
	Artistlist    *[]Artist
	SortingFilter string
	Asc           bool
	*Page
	IdPageToDisplay int
	*ListenAddr
}

/*Spotify struct*/

type SpotifyToken struct {
	Access_token string
	Token_type   string
	Expires_in   int
}
