package main

import "fmt"
import "github.com/PotatoesInHD/pokedex/internal/pokeapi"
import "math/rand"

func commandCatch(cfg *config, parameter string) error {
	if parameter == "" {
		fmt.Println("catch <pokemon> can't be empty")
		return nil
	}
	url := "https://pokeapi.co/api/v2/pokemon/" + parameter
	catchRes, err := pokeapi.CatchPokemon(url, cfg.catchCachePtr)
	if err != nil {
		fmt.Printf("Error after CatchPokemon in commandCatch: %v\n", err)
		return err
	}
	name := catchRes.Name
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	random := rand.Intn(catchRes.BaseExperience * 3)
	difficulty := catchRes.BaseExperience
	fmt.Printf("%d %d\n", difficulty, random)
	if cfg.collectedPokes == nil {
		cfg.collectedPokes = make(map[string]pokeapi.CatchResponse)
	}
	if random > difficulty {
		fmt.Printf("Caught Pokemon: %s!\n", name)
		cfg.collectedPokes[name] = catchRes
	} else {
		fmt.Printf("Missed Pokemon: %s :(\n", name)
	}
	return nil
}
