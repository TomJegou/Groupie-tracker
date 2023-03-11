package api

import (
	"absolut-music/src/globalDataStructures"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

func PreprocessArtNameSearchSpotify(artistName string) string {
	result := ""
	l := strings.Split(artistName, " ")
	for index, kword := range l {
		if index != len(l)-1 {
			kword += "%20"
		}
		result += kword
	}
	return result
}

func MakeReqSearchAPISportify(artistName string) *http.Request {
	url := "https://api.spotify.com/v1/search?q=" + PreprocessArtNameSearchSpotify(artistName) + "&type=artist&offset=0&limit=1"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	PutBodyResponseApiIntoStruct(RequestApi(MakeReqTokenSpotify()), &globalDataStructures.OAuthSpotifyToken, &wg)
	wg.Wait()
	req.Header.Set("Authorization", "Bearer "+globalDataStructures.OAuthSpotifyToken.Access_token)
	return req
}

func MakeReqTokenSpotify() *http.Request {
	data := url.Values{"grant_type": {"client_credentials"}}
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", "Basic "+globalDataStructures.EncodedAuth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.PostForm = data
	return req
}
