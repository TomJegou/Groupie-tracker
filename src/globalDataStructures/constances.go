package globalDataStructures

import (
	"absolut-music/src/structures"
	"text/template"
)

/*Global constances*/

const URLARTISTS = "https://groupietrackers.herokuapp.com/api/artists"
const URLDATES = "https://groupietrackers.herokuapp.com/api/dates"
const URLLOCATIONS = "https://groupietrackers.herokuapp.com/api/locations"
const URLRELATION = "https://groupietrackers.herokuapp.com/api/relation"

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

/*Global channels*/

var ChanArtists = make(chan *[]structures.Artist)
var ChanTemplates = make(chan *template.Template, 1)
var ChanArtDet = make(chan *structures.Artist)
