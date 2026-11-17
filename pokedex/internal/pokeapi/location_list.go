package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"github.com/PotatoesInHD/pokedex/internal/pokecache"
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


func GetAreas(url string, cache *pokecache.Cache) (AreaResponse, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	//checks to see if cache exists. if so unmarshal if not do a new http request
	if bodyCache, ok := cache.GetCache(url); ok {
		areas := AreaResponse{}
		if err := json.Unmarshal(bodyCache, &areas); err != nil {
		    return AreaResponse{}, err
		}
		return areas, nil
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

	cache.AddCache(url, body)
	areas := AreaResponse{}
	if err = json.Unmarshal(body, &areas); err != nil {
	    return AreaResponse{}, err
	}
	return areas, nil
}
