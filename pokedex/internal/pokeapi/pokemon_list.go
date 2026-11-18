package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"github.com/PotatoesInHD/pokedex/internal/pokecache"
)

type FindResponse struct {
	Id int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	GameIndex int `json:"game_index,omitempty"`
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters,omitempty"`
}
type Pokemon struct {
	Name string `json:"name,omitempty"`
	Url string `json:"url,omitempty"`
}
type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon,omitempty"`
}


func FindPokemon(url string, cache *pokecache.Cache) (FindResponse, error) {
	//checks to see if cache exists. if so unmarshal if not do a new http request
	if bodyCache, ok := cache.GetCache(url); ok {
		foundPokemon := FindResponse{}
		if err := json.Unmarshal(bodyCache, &foundPokemon); err != nil {
		    return FindResponse{}, err
		}
		return foundPokemon, nil
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
	foundPokemon := FindResponse{}
	if err = json.Unmarshal(body, &foundPokemon); err != nil {
	    return FindResponse{}, err
	}
	return foundPokemon, nil
}
