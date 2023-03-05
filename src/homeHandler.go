package src

import (
	"fmt"
	"net/http"
)

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

/*Home page's handler*/
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(GetIP(r))
	OnLibraryArtists = false
	go ParseHtml("static/html/index.html")
	template := <-ChanTemplates
	template.Execute(w, ListeningAddr)
}
