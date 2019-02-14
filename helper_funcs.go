package newsapi

import (
	"log"
	"net/http"
	"reflect"
	"time"
)

func checkStringInArray(s string, sl []string) bool {
	for _, e := range sl {
		if e == s {
			return true
		}
	}
	return false
}

func makeRequest(url string, payload jsonPayload, apiKey string) (*http.Response, error) {

	client := &http.Client{Timeout:time.Duration(30 * time.Second)}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Could not create the request")
	}
	req.Header.Add("x-api-key", apiKey) // set api key in the header for authentication

	// Build query string
	q := req.URL.Query()
	v := reflect.ValueOf(payload)

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).String() != "" {  // ignore any empty params
			q.Add(reflect.Indirect(reflect.ValueOf(payload)).Type().Field(i).Name, v.Field(i).String())
		}
	}
	req.URL.RawQuery = q.Encode()

	return client.Do(req)
}