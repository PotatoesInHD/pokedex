package main

import (
	"fmt"
	"time"

	"github.com/PotatoesInHD/pokedex/internal/pokeapi"
	"github.com/PotatoesInHD/pokedex/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(15 * time.Second)
	areaRes, err := pokeapi.GetAreas("", cache)
	if err != nil {
		fmt.Printf("Error after get areas in Main: %v \n", err)
	}
	commands := getCommands()
	cfg := &config{commands: commands, nextLocUrl: areaRes.Next, prevLocUrl: areaRes.Prev, cache: cache}
	cfg.nextLocUrl = nil
	startRepl(cfg)
}
