package globalDataStructures

import (
	"absolut-music/src/structures"
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
var SearchContent string
var SortingOption string
var SortingOrder string
var PaginationRequest string
var NumberOfElem string
