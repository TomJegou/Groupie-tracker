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

/*Search struct*/

type SpotifySearchImage struct {
	Height int
	Url    string
	Width  int
}

/*Search Artist*/
type SpotifySearchArtist struct {
	Artists SpotifyResultSearchArtist
}

type SpotifyResultSearchArtist struct {
	Hrefs  string
	Items  []SpotifyResultSearchArtItem
	Limit  int
	Next   string
	Offset int
	Total  int
}

type SpotifyResultSearchArtItem struct {
	External_urls map[string]string
	Followers     SpotifyResultSearchArtFollowers
	Genres        []string
	Href          string
	Id            string
	Images        []SpotifySearchImage
	Name          string
	Popularity    int
	Type          string
	Uri           string
}

type SpotifyResultSearchArtFollowers struct {
	Href  string
	Total int
}

/*Search Album*/
type SpotifySearchAlbum struct {
	Albums SpotifyResultSearchAlbum
}

type SpotifyResultSearchAlbum struct {
	Href     string
	Items    []SpotifyResultSearchAlbumItem
	Limit    int
	Next     string
	Offset   int
	Previous interface{}
	Total    int
}

type SpotifyResultSearchAlbumItem struct {
	Album_group            string
	Album_type             string
	Artists                []SpotifyResultSearchAlbumArtist
	Available_markets      []string
	External_urls          map[string]string
	Href                   string
	Id                     string
	Images                 []SpotifySearchImage
	Is_playable            bool
	Name                   string
	Release_date           string
	Release_date_precision string
	Total_tracks           int
	Type                   string
	Uri                    string
}

type SpotifyResultSearchAlbumArtist struct {
	External_urls map[string]string
	Href          string
	Id            string
	Name          string
	Type          string
	Uri           string
}
