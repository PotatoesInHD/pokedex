package main

import (
	"time"
	"github.com/PotatoesInHD/pokedex/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Minute)
	commands := getCommands()
	cfg := &config{commands: commands, nextLocUrl: nil, prevLocUrl: nil, cache: cache}
	startRepl(cfg)
}
