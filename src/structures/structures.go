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
	*SpotifySearchArtist
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

/*tools*/

type FormatDate struct {
	Year  int
	Month int
	Day   int
}

/*Spotify struct*/

type SpotifyToken struct {
	Access_token string
	Token_type   string
	Expires_in   int
}

type SpotifySearchArtist struct {
	Artists SpotifyArtist
}

type SpotifyArtist struct {
	Hrefs  string
	Items  []SpotifyArtItem
	Limit  int
	Next   string
	Offset int
	Total  int
}

type SpotifyArtItem struct {
	External_urls map[string]string
	Followers     SpotifyArtFollowers
	Genres        []string
	Href          string
	Id            string
	Images        []SpotifyArtImage
	Name          string
	Popularity    int
	Type          string
	Uri           string
}

type SpotifyArtFollowers struct {
	Href  string
	Total int
}

type SpotifyArtImage struct {
	Height int
	Url    string
	Width  int
}
