package src

import (
	"net/http"
	"text/template"
)

func LegalNoticeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("static/html/legalNotice.html")
	template.Execute(w, nil)
}
