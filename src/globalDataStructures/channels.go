package globalDataStructures

import (
	"absolut-music/src/structures"
	"text/template"
)

/*Global channels*/

var ChanArtists = make(chan *[]structures.Artist)
var ChanTemplates = make(chan *template.Template, 1)
var ChanArtDet = make(chan *structures.Artist)
