package main

import (
	"time"
	"github.com/PotatoesInHD/pokedex/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Minute)
	commands := getCommands()
	cfg := &config{commands: commands, cachePtr: cache}
	startRepl(cfg)
}
