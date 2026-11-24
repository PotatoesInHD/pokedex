package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/PotatoesInHD/pokedex/internal/pokecache"
	"github.com/PotatoesInHD/pokedex/internal/pokeapi"
)


type cliCommand struct {
	name 		string
	description string
	callback 	func(cfg *config, parameter string) error
}

type config struct {
	commands 		map[string]cliCommand
	nextLocUrl		*string
	prevLocUrl		*string
	areaCachePtr 	*pokecache.Cache
	pokemonCachePtr *pokecache.Cache
	catchCachePtr 	*pokecache.Cache
	collectedPokes  map[string]pokeapi.CatchResponse
}


func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %v", err)
	}
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		result := cleanInput(scanner.Text())
		var parameter string
		if len(result) >= 2 {
			 parameter = result[1]
		}
		if len(result) != 0 {
			if command, ok := cfg.commands[result[0]]; ok {
				if err := command.callback(cfg, parameter); err != nil {
					fmt.Printf("Error in startRepl: %v", err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}


func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex. Ctrl+C will also exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays next 20 locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays previous 20 locations",
			callback: commandMapb,
		},
		"explore": {
			name: "explore <destination>",
			description: "Displays the names of all Pokemon in the area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch <pokemon>",
			description: "Attempts to catch pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect <pokemon>",
			description: "Inspects captured pokemon",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Lists all captured pokemon",
			callback: commandPokedex,
		},
	}
	return commands
}


func cleanInput(text string) []string {
	cleanSlice := strings.Fields(strings.ToLower(text))
	return  cleanSlice
}
