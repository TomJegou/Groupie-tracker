package src

import (
	"fmt"
	"sync"
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

func RunParallel(f func(*sync.WaitGroup)) {
	var wg sync.WaitGroup
	wg.Add(1)
	go f(&wg)
	wg.Wait()
}
