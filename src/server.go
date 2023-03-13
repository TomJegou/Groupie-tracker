package src

import (
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/handlers"
	"fmt"
	"net/http"
	"sync"
)

/*
Establish the routing for the webApp and start the server
on port 80
*/
func StartServer(wg *sync.WaitGroup) {
	defer wg.Done()
	FileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", FileServer))
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/libraryArtists", libraryArtistsHandler)
	http.HandleFunc("/artistsDetails", ArtistsDetailsHandler)
	http.HandleFunc("/about", AboutHandler)
	http.HandleFunc("/legalNotice", LegalNoticeHandler)
	http.HandleFunc("/historic", HistoricHandler)
	fmt.Println("http://localhost:80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error starting the server")
	}

}
