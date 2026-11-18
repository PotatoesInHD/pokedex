package main

import "fmt"
import "os"


func commandExit(cfg *config, parameter string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
