package structures

type SpotifyToken struct {
	Access_token string
	Token_type   string
	Expires_in   int
}

/*Search struct*/

type SpotifyImage struct {
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
	Images        []SpotifyImage
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
	Images                 []SpotifyImage
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

/*Spotify artist/id/albums*/

type SpotifyArtistAlbums struct {
	Href     string
	Items    []SpotifyArtistAlbumsItems
	Limit    int
	Next     string
	Offset   int
	Previous string
	Total    int
}

type SpotifyArtistAlbumsItems struct {
	Album_type             string
	Total_tracks           int
	Available_markets      []string
	External_urls          map[string]string
	Href                   string
	Id                     string
	Images                 []SpotifyImage
	Name                   string
	Release_date           string
	Release_date_precision string
	Type                   string
	Uri                    string
	Genres                 []string
	Popularity             int
	Artists                []SpotifyArtistAlbumsArtists
}

type SpotifyArtistAlbumsArtists struct {
	External_urls map[string]string
	Href          string
	Id            string
	Name          string
	Type          string
	Uri           string
}
