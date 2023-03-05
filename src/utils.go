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

func ChangeListenAddr(r *http.Request) {
	if r.Host != ListeningAddr.Ipv4 {
		ListeningAddr.Ipv4 = r.Host
	}
}
