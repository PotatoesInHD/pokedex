package main

import "fmt"


func commandPokedex(cfg *config, parameter string) error {
	if len(cfg.collectedPokes) == 0 {
		fmt.Println("You're Pokedex is empty")
	} else {
	fmt.Println("You're Pokedex:")
	for _, pokemon := range cfg.collectedPokes {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	}
	return nil
}
