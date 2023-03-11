package api

import (
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

func MakeReqSearchAPISportify(artistName string) *http.Request {
	url := "https://api.spotify.com/v1/search?q=" + tools.PreprocessArtNameSearchSpotify(artistName) + "&type=artist&offset=0&limit=1"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	PutBodyResponseApiIntoStruct(RequestApi(MakeReqTokenSpotify()), &gds.OAuthSpotifyToken, &wg)
	wg.Wait()
	req.Header.Set("Authorization", "Bearer "+gds.OAuthSpotifyToken.Access_token)
	return req
}

func MakeReqTokenSpotify() *http.Request {
	data := url.Values{"grant_type": {"client_credentials"}}
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", "Basic "+gds.EncodedAuth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.PostForm = data
	return req
}
