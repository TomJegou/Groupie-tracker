package src

import (
	"net/http"
	"text/template"
	"fmt"
)

func AboutHandlerFunc(w http.ResponseWriter, r *http.Request) {
	template, errors := template.ParseFiles("static/html/about.html")
	if errors != nil {
		fmt.Println("Error Parsing Template")
		fmt.Println(errors)
	}
	template.Execute(w, nil)
}
