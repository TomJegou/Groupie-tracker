package handlers

import (
	"net/http"
)

func AddHeaderFs(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		fs.ServeHTTP(w, r)
	}
}
