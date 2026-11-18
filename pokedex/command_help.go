package main

import "fmt"
import "strings"
import "sort"

func commandHelp(cfg *config, parameter string) error {
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	commandNames := []string{}
	for name := range cfg.commands {
		commandNames = append(commandNames, name)
	}
	sort.Strings(commandNames)
	for _, commandName := range commandNames {
		fmt.Printf("%s - %s\n", commandName, cfg.commands[commandName].description)
	}
	fmt.Println(strings.Repeat("-", 30))
	return nil
}
