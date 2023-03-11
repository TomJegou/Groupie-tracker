package api

import (
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"fmt"
	"net/http"
	u "net/url"
	"strings"
	"sync"
)

/*Create a request with the headers and the forms and return the request's pointers*/
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

/*Create a request for the Herokuapp api and return the request's pointer*/
func MakeReqHerokuapp(url string) *http.Request {
	return MakeReq(url, nil, nil)
}

/*
Create a request to get the result of Spotify research of the artist name
and return the request's pointer
*/
func MakeReqSearchArtAPISportify(artistName string) *http.Request {
	url := "https://api.spotify.com/v1/search?q=" + tools.PreprocessArtNameSearchSpotify(artistName) + "&type=artist&offset=0&limit=1"
	var wg sync.WaitGroup
	wg.Add(1)
	go PutBodyResponseApiIntoStruct(RequestApi(MakeReqTokenSpotify()), &gds.OAuthSpotifyToken, &wg)
	wg.Wait()
	return MakeReq(url, map[string]string{"Authorization": "Bearer " + gds.OAuthSpotifyToken.Access_token}, nil)
}

/*Create a request to get the Spotify Authorisation token and return the request's pointer*/
func MakeReqTokenSpotify() *http.Request {
	headers := map[string]string{"Authorization": "Basic " + gds.EncodedAuth, "Content-Type": "application/x-www-form-urlencoded"}
	forms := map[string]string{"grant_type": "client_credentials"}
	return MakeReq("https://accounts.spotify.com/api/token", headers, forms)
}

/*
Create a request to get the result of Spotify research of the artist name
and return the request's pointer
*/
func MakeReqSearchAlbumArtAPISportify(artistName string) *http.Request {
	url := "https://api.spotify.com/v1/search?q=" + tools.PreprocessArtNameSearchSpotify(artistName) + "&type=album&offset=0&limit=1"
	var wg sync.WaitGroup
	wg.Add(1)
	go PutBodyResponseApiIntoStruct(RequestApi(MakeReqTokenSpotify()), &gds.OAuthSpotifyToken, &wg)
	wg.Wait()
	return MakeReq(url, map[string]string{"Authorization": "Bearer " + gds.OAuthSpotifyToken.Access_token}, nil)
}

/*
Create a request to get the result of Spotify research of the artist name
and return the request's pointer
*/
func MakeReqArtAlbumAPISportify(artistId string) *http.Request {
	url := "https://api.spotify.com/v1/artists/" + artistId + "/albums"
	var wg sync.WaitGroup
	wg.Add(1)
	go PutBodyResponseApiIntoStruct(RequestApi(MakeReqTokenSpotify()), &gds.OAuthSpotifyToken, &wg)
	wg.Wait()
	return MakeReq(url, map[string]string{"Authorization": "Bearer " + gds.OAuthSpotifyToken.Access_token}, nil)
}
