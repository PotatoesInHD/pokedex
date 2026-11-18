package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"github.com/PotatoesInHD/pokedex/internal/pokecache"
)


type CatchResponse struct {
	Id int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	BaseExperience int `json:"base_experience,omitempty"`
	Height int `json:"height,omitempty"`
	Weight int `json:"weight,omitempty"`
	Abilities []Abilities `json:"abilities,omitempty"`
	Stats []Stats `json:"stats,omitempty"`
	Types []Types `json:"types,omitempty"`
}
type Ability struct {
	Name string `json:"name,omitempty"`
	Url string `json:"url,omitempty"`
}
type Abilities struct {
	IsHidden bool `json:"is_hidden,omitempty"`
	Slot int `json:"slot,omitempty"`
	Ability Ability `json:"ability,omitempty"`
}
type Stat struct {
	Name string `json:"name,omitempty"`
	Url string `json:"url,omitempty"`
}
type Stats struct {
	BaseStat int `json:"base_stat,omitempty"`
	Effort int `json:"effort,omitempty"`
	Stat Stat `json:"stat,omitempty"`
}
type Type struct {
	Name string `json:"name,omitempty"`
	Url string `json:"url,omitempty"`
}
type Types struct {
	Slot int `json:"slot,omitempty"`
	Type Type `json:"type,omitempty"`
}


func CatchPokemon(url string, cache *pokecache.Cache) (CatchResponse, error) {
	//checks to see if cache exists. if so unmarshal if not do a new http request
	if bodyCache, ok := cache.GetCache(url); ok {
		catchPokemon := CatchResponse{}
		if err := json.Unmarshal(bodyCache, &catchPokemon); err != nil {
		    return CatchResponse{}, err
		}
		return catchPokemon, nil
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
	catchPokemon := CatchResponse{}
	if err = json.Unmarshal(body, &catchPokemon); err != nil {
	    return CatchResponse{}, err
	}
	return catchPokemon, nil
}
