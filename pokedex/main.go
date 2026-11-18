package main

import (
	"time"
	"github.com/PotatoesInHD/pokedex/internal/pokecache"
)

func main() {
	areaCache := pokecache.NewCache(5 * time.Minute)
	pokemonCache := pokecache.NewCache(5 * time.Minute)
	commands := getCommands()
	cfg := &config{commands: commands, areaCachePtr: areaCache, pokemonCachePtr: pokemonCache}
	startRepl(cfg)
}
