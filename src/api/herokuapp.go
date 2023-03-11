package api

import (
	"fmt"
	"net/http"
)

func MakeReqHerokuapp(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	return req
}
