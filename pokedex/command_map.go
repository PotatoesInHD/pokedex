package main

import "github.com/PotatoesInHD/pokedex/internal/pokeapi"
import "fmt"


func commandMap(cfg *config, parameter string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.nextLocUrl != nil {
		url = *cfg.nextLocUrl
	}
	areaRes, err := pokeapi.GetAreas(url, cfg.areaCachePtr)
	if err != nil {
		fmt.Printf("Error GetAreas in commandMap: %v\n", err)
		return err
	}
	cfg.nextLocUrl = areaRes.Next
	cfg.prevLocUrl = areaRes.Prev
	for i := 0; i < len(areaRes.Results); i++{
		fmt.Println(areaRes.Results[i].Name)
	}
	return nil
}
