package api

import (
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/tools"
	"net/http"
	"sync"
)

func MakeReqSearchAPISportify(artistName string) *http.Request {
	url := "https://api.spotify.com/v1/search?q=" + tools.PreprocessArtNameSearchSpotify(artistName) + "&type=artist&offset=0&limit=1"
	var wg sync.WaitGroup
	wg.Add(1)
	PutBodyResponseApiIntoStruct(RequestApi(MakeReqTokenSpotify()), &gds.OAuthSpotifyToken, &wg)
	wg.Wait()
	return MakeReq(url, map[string]string{"Authorization": "Bearer " + gds.OAuthSpotifyToken.Access_token}, nil)
}

func MakeReqTokenSpotify() *http.Request {
	headers := map[string]string{"Authorization": "Basic " + gds.EncodedAuth, "Content-Type": "application/x-www-form-urlencoded"}
	forms := map[string]string{"grant_type": "client_credentials"}
	return MakeReq("https://accounts.spotify.com/api/token", headers, forms)
}
