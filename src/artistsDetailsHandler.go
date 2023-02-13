package src

import (
	"fmt"
	"net/http"
	"text/template"
)

func ArtistsDetailsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("artistCard"))
	template, _ := template.ParseFiles("static/html/artistsDetails.html")
	template.Execute(w, nil)
}
