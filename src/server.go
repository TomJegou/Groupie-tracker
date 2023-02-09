package src

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Accueil(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("static/html/index.html")
	template.Execute(w, nil)
}

func StartServer() {
	FileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", FileServer))
	http.HandleFunc("/", Accueil)
	http.HandleFunc("/artists", ArtistsHandlerFunc)
	http.HandleFunc("/artistsDetails", ArtistsDetailsHandlerFunc)
	http.HandleFunc("/about", AboutHandlerFunc)
	http.HandleFunc("/legalNotice", LegalNoticeHandlerFunc)
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
