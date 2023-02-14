package src

import (
	"fmt"
	"net/http"
	"text/template"
)

func LegalNoticeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	template, errors := template.ParseFiles("static/html/legalNotice.html")
	if errors != nil {
		fmt.Println("Error Parsing Template")
	}
	template.Execute(w, nil)
}
