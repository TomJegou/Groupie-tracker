package api

import (
	"fmt"
	"net/http"
	u "net/url"
	"strings"
)

func MakeReq(url string, headers map[string]string, forms map[string]string) *http.Request {
	if forms != nil {
		data := u.Values{}
		for k, v := range forms {
			data.Add(k, v)
		}
		req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
		if err != nil {
			fmt.Println(err)
		}
		req.PostForm = data
		for k, v := range headers {
			req.Header.Set(k, v)
		}
		return req
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err)
		}
		for k, v := range headers {
			req.Header.Set(k, v)
		}
		return req
	}
}
