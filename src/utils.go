package src

import (
	"fmt"
	"net/http"
	"sync"
	"text/template"
)

/*
Parse the html file passed as a parameter and send the
template into  the ChanTemplates Channel
*/
func ParseHtml(fileToParse string) {
	template, errors := template.ParseFiles(fileToParse)
	if errors != nil {
		fmt.Println("Error Parsing Template")
		fmt.Println(errors)
	}
	ChanTemplates <- template
}

/*Make the function passed as a parameter run in Parallel as a goroutine*/
func RunParallel(f func(*sync.WaitGroup)) {
	var wg sync.WaitGroup
	wg.Add(1)
	go f(&wg)
	wg.Wait()
}

/*
Check if in the request if the host ipv4 is
the same as the one to be used for the templates.
If it's not the same, change the ListeningAddr.Ipv4 to the host requested
*/
func ChangeListenAddr(r *http.Request) {
	if r.Host != ListeningAddr.Ipv4 {
		ListeningAddr.Ipv4 = r.Host
	}
}
