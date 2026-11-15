package main
import "fmt"
import "strings"

func commandHelp(cfg *config) error {
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println(strings.Repeat("-", 30))
	return nil
}
