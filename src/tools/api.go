package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

/*Do an API call and return a string of the response*/
func GetApi(url string) []byte {
	req, errors := http.NewRequest("GET", url, nil)
	if errors != nil {
		fmt.Println("Error Request")
		fmt.Println(errors)
		return []byte{}
	}
	res, errors := http.DefaultClient.Do(req)
	if errors != nil {
		fmt.Println("Error default client")
		fmt.Println(errors)
		return []byte{}
	}
	defer res.Body.Close()
	body, errors := io.ReadAll(res.Body)
	if errors != nil {
		fmt.Println("Error during read body")
		fmt.Println(errors)
		return []byte{}
	}
	return body
}

/*
Call the API using the url passed as a parameter
and the func GetApi, and put the response into the structure passed as a parameter
*/
func PutBodyResponseApiIntoStruct(url string, structure interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	err := json.Unmarshal(GetApi(url), &structure)
	if err != nil {
		fmt.Println("Erreur Unmarshal JSON\n", err)
	}
}
