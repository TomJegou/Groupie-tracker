package api

import (
	"net/http"
)

func MakeReqHerokuapp(url string) *http.Request {
	return MakeReq(url, nil, nil)
}
