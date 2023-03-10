package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

/*Do an API call and return the body of the response*/
func RequestApi(req *http.Request) []byte {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body
}

/*
Call the API using the url passed as a parameter
and the func GetApi, and put the response into the structure passed as a parameter
*/
func PutBodyResponseApiIntoStruct(body []byte, structure interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	err := json.Unmarshal(body, &structure)
	if err != nil {
		fmt.Println("Erreur Unmarshal JSON\n", err)
	}
}
