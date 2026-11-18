package main

import (
	"time"
	"github.com/PotatoesInHD/pokedex/internal/pokecache"
	"fmt"
)

func main() {
	fmt.Println("Type: 'help' for list of commands")
	areaCache := pokecache.NewCache(5 * time.Minute)
	pokemonCache := pokecache.NewCache(5 * time.Minute)
	catchCache := pokecache.NewCache(5 * time.Minute)
	commands := getCommands()
	cfg := &config{commands: commands, areaCachePtr: areaCache, pokemonCachePtr: pokemonCache, catchCachePtr: catchCache}
	startRepl(cfg)
}
