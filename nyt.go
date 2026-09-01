package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type nytResponse struct {
	Solution string `json:"solution"`
}

func fetchNYTWord() string {
	datestr := time.Now().Format("2006-01-02")
	url := "https://www.nytimes.com/svc/wordle/v2/" + datestr + ".json"
	resp, err := http.Get(url)
	if err != nil {
		panic("error while fetching nyt solution. Error: " + err.Error())
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			panic("error closing response body: " + closeErr.Error())
		}
	}()
	var response nytResponse

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("error while reading nyt solution. Error: " + err.Error())
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		panic("error while parsing nyt solution. Error: " + err.Error())
	}

	return response.Solution
}
