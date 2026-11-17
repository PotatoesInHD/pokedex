package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


type cliCommand struct {
	name string
	description string
	callback func(cfg *config) error
}

type config struct {
	commands map[string]cliCommand
	nextLocUrl *string
	prevLocUrl *string
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
		if len(result) != 0 {
			if command, ok := cfg.commands[result[0]]; ok {
				if err := command.callback(cfg); err != nil {
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
			description: "Exit the Pokedex",
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
	}
	return commands
}


func cleanInput(text string) []string {
	cleanSlice := []string{}
	cleanSlice = strings.Fields(strings.ToLower(text))
	return  cleanSlice
}
