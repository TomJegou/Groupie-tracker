package src

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type DateLocationConcertRelation struct {
	*Artist
}

func ArtistsDetailsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	idInt, err := strconv.Atoi(r.FormValue("artistCard"))
	if err != nil {
		fmt.Println("Error converting string to integer")
		fmt.Println(err)
	}
	template, err := template.ParseFiles("static/html/artistsDetails.html")
	if err != nil {
		fmt.Println("Error parsing template artistsDetails.html")
		fmt.Println(err)
	}
	artist := Artists[idInt]
	datelocationmachintruc := DateLocationConcertRelation{}
	datelocationmachintruc.Artist = &artist
	template.Execute(w, datelocationmachintruc)
}
