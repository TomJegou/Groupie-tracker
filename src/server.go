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
func StartServer(wg *sync.WaitGroup, startAttempt int) {
	// Check if server failed to start and print error message if so
	if startAttempt <= 0 {
		fmt.Printf("Error: tried start the serveur %v, but failed\n Need to restart the server manually.", startAttempt)
	} else {
		// Decrease the WaitGroup counter when the function returns
		defer wg.Done()
		// Set up file server for static files
		fileServer := http.FileServer(http.Dir("./static"))
		http.Handle("/static/", handlers.AddHeaderFs(http.StripPrefix("/static", fileServer)))
		// Set up routing for different pages
		http.HandleFunc("/", handlers.HomeHandler)
		http.HandleFunc("/libraryArtists", handlers.LibraryArtistsHandler)
		http.HandleFunc("/artistsDetails", handlers.ArtistsDetailsHandler)
		http.HandleFunc("/about", handlers.AboutHandler)
		http.HandleFunc("/legalNotice", handlers.LegalNoticeHandler)
		http.HandleFunc("/location", handlers.LocationHandler)
		http.HandleFunc("/spotify-search", handlers.SpotifyHandler)
		http.HandleFunc("/album-detail", handlers.AlbumDetHandler)
		http.HandleFunc("/concertDet", handlers.ConcertDetailHandler)
		http.HandleFunc("/date-historic", handlers.DateHistoricHandler)
		// Start the server
		fmt.Println("http://127.0.0.1:80")
		err := http.ListenAndServe(gds.ListeningAddr.Ipv4+":"+gds.ListeningAddr.Port, nil)
		if err != nil {
			// If server failed to start, print error message and attempt to start again
			fmt.Println(err)
			fmt.Println("Error starting the server")
			StartServer(wg, startAttempt-1)
		}
	}

}
