package main

import "fmt"
import "github.com/PotatoesInHD/pokedex/internal/pokeapi"


func commandExplore(cfg *config, parameter string) error {
	if parameter == "" {
		fmt.Println("explore <destination> can't be empty")
		return nil
	}

	url := "https://pokeapi.co/api/v2/location-area/" + parameter
	findRes, err := pokeapi.FindPokemon(url, cfg.pokemonCachePtr)
	if err != nil {
		fmt.Printf("Error after FindPokemon in commandExplore: %v\n", err)
		return err
	}
	fmt.Printf("Exploring %s...\n", findRes.Name)
	fmt.Println("Found Pokemon:")
	for i := 0; i < len(findRes.PokemonEncounters); i++{
		fmt.Println(findRes.PokemonEncounters[i].Pokemon.Name)
	}
	return nil
}
