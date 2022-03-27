package main

import (
	"log"
	"net/http"
)

// This function makes an http GET request
func httpGet(url string) *http.Response {

	log.Println("Making GET request to: ", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("There was an error with the HTTP request: ", err)
		return &http.Response{}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("There was an error with the HTTP response: ", err)
		return &http.Response{}
	}

	return res
}
