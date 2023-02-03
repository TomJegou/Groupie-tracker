package src

import (
	"log"
	"net/http"
	"text/template"
)

func Accueil(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("static/html/index.html")
	template.Execute(w, nil)
}

func StartServer() {
	http.HandleFunc("/", Accueil)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
