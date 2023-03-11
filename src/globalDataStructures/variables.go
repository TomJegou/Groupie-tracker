package globalDataStructures

import (
	"absolut-music/src/structures"
	b64 "encoding/base64"
)

/*Global variables*/
var ListeningAddr = structures.ListenAddr{Ipv4: "0.0.0.0", Port: "80"}
var IsStartServer = true
var OnLibraryArtists = false
var Artists []structures.Artist
var Dates map[string][]structures.Date
var Locations map[string][]structures.Location
var Relations map[string][]structures.Relation
var PageCapacity int
var LibArtists structures.LibraryArtists
var ListPages []structures.Page
var SearchContent = ""
var SortingOption = ""
var SortingOrder = ""
var PaginationRequest = ""
var NumberOfElem = ""
var EncodedAuth = b64.StdEncoding.EncodeToString([]byte(SPOTIFY_CLIENT_ID + ":" + SPOTIFY_CLIENT_SECRET))
var OAuthSpotifyToken = &structures.SpotifyToken{}
var ResultSpotifySearchArtist = &structures.SpotifySearchArtist{}
var ResultSpotifySearchAlbum = &structures.SpotifySearchAlbum{}
var ResultSpotifyArtistAlbums = &structures.SpotifyArtistAlbums{}
