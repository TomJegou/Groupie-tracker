package tools

import (
	"absolut-music/src/globalDataStructures"
	"absolut-music/src/structures"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func SearchAPISportify(artistName string) string {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/search?q="+artistName+"&type=artist", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", "Bearer "+GetTokenSpotify())
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func GetTokenSpotify() string {
	data := url.Values{"grant_type": {"client_credentials"}}
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", "Basic "+globalDataStructures.EncodedAuth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.PostForm = data
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var token = &structures.SpotifyToken{}
	err = json.Unmarshal(body, token)
	if err != nil {
		fmt.Println("Erreur Unmarshal JSON\n", err)
	}
	return token.Access_token
}
