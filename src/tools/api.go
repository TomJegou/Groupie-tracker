package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

/*Functions*/

/*Do an API call and return a string of the response*/
func GetApi(url string) string {
	req, errors := http.NewRequest("GET", url, nil)
	if errors != nil {
		fmt.Println("Error Request")
		fmt.Println(errors)
		return ""
	}
	res, errors := http.DefaultClient.Do(req)
	if errors != nil {
		fmt.Println("Error default client")
		fmt.Println(errors)
		return ""
	}
	defer res.Body.Close()
	body, errors := io.ReadAll(res.Body)
	if errors != nil {
		fmt.Println("Error during read body")
		fmt.Println(errors)
		return ""
	}
	return string(body)
}

/*
Call the API using the url passed as a parameter
and the func GetApi, and put the response into the structure passed as a parameter
*/
func PutBodyResponseApiIntoStruct(url string, structure interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	err := json.Unmarshal([]byte(GetApi(url)), &structure)
	if err != nil {
		fmt.Println("Erreur Unmarshal JSON\n", err)
	}
}