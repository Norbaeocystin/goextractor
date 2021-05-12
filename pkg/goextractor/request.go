package goextractor

import (
	"io/ioutil"
	"net/http"
	"time"
)

var HeadersDefault = map[string][]string{
	"User-Agent":      {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36 Edg/85.0.564.51"},
	"Accept-Language": {"en-US"},
}

// Change if you want
const timeout = 20

//Send get request method. Returns bytes and error, hard coded timeout 20 secods
func GetWithHeaders(urlstring string, headers map[string][]string) ([]byte, error) {
	client := http.Client{Timeout: timeout * time.Second}
	request, err := http.NewRequest("GET", urlstring, nil)
	if err != nil {
		return []byte(""), err
	}
	request.Header = headers
	// Make request
	response, err := client.Do(request)
	if err != nil {
		return []byte(""), err
	}
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return bodyBytes, err
}
