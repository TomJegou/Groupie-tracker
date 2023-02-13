package src

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func ArtistsDetailsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	idInt, err := strconv.Atoi(r.FormValue("artistCard"))
	if err != nil {
		fmt.Println(err)
	}
	var artistToDisplay = Artists[idInt]
	template, _ := template.ParseFiles("static/html/artistsDetails.html")
	template.Execute(w, artistToDisplay)
}
