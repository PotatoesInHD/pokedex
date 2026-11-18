package main

import "fmt"


func commandInspect(cfg *config, parameter string) error {
	if parameter == "" {
		fmt.Println("inspect <pokemon> can't be empty")
		return nil
	}

	if pokeInfo, ok := cfg.collectedPokes[parameter]; ok {
		fmt.Printf("Name: %s\n", pokeInfo.Name)
		fmt.Printf("Height: %d\n", pokeInfo.Height)
		fmt.Printf("Weight: %d\n", pokeInfo.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokeInfo.Stats {
			fmt.Printf("    %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		for _, pokeType := range pokeInfo.Types{
			fmt.Printf("    - %s\n", pokeType.Type.Name)
		}
		return nil
	} else {
		fmt.Printf("%s not captured yet\n", parameter)
		return nil
	}
}
