package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)


type AreaResponse struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Prev 	 *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


func GetAreas(url string) (AreaResponse, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error during Get: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatalf("io.ReadAll resonse error: %v", err)
	}

	areas := AreaResponse{}
	if err = json.Unmarshal(body, &areas); err != nil {
	    return AreaResponse{}, err
	}
	return areas, nil
}
