package structures

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
	Albums *[]SpotifyArtistAlbumsItems
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

/*Structure used for the Spotify page*/

type SpotifyHandlerStruct struct {
	*ListenAddr
	*SpotifySearchArtist
}

/*Structure used for the Album details page*/
type AlbumDetail struct {
	*ListenAddr
	*SpotifyAlbum
}

/*tools*/

type FormatDate struct {
	Year  int
	Month int
	Day   int
}
