package main

import (
	"fmt"
	"github.com/PotatoesInHD/pokedex/internal/pokeapi"
)

func main() {
	areaRes, err := pokeapi.GetAreas("")
	if err != nil {
		fmt.Printf("Error after get areas in Main: %v \n", err)
	}
	commands := getCommands()
	cfg := &config{commands: commands, nextLocUrl: areaRes.Next, prevLocUrl: areaRes.Prev}
	cfg.nextLocUrl = nil
	startRepl(cfg)
}
