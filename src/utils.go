package src

import (
	"fmt"
	"text/template"
)

func ParseHtml(fileToParse string) {
	template, errors := template.ParseFiles(fileToParse)
	if errors != nil {
		fmt.Println("Error Parsing Template")
		fmt.Println(errors)
	}
	ChanTemplates <- template
}
