package handlers

import (
	"fmt"
	"net/http"
)

func JsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint static/scripts")
	w.Header().Set("Content-Type", "text/javascript")
	http.ServeFile(w, r, "static/scripts")
}
