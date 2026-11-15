package main
import "fmt"
import "strings"

func commandHelp(cfg *config) error {
	//order changes everytime. maybe sweeten later by putting
	// into a slice, sorting and then looking up/printing
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, command := range cfg.commands {
		fmt.Printf("%s - %s\n", command.name, command.description)
	}
	fmt.Println(strings.Repeat("-", 30))
	return nil
}
