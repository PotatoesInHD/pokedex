package main
import "strings"
import "os"
import "bufio"
import "fmt"


type cliCommand struct {
	name string
	description string
	callback func(cfg *config) error
}

type config struct {
	commands map[string]cliCommand
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		result := cleanInput(scanner.Text())
		if len(result) != 0 {
			if command, ok := cfg.commands[result[0]]; ok {
				if err := command.callback(cfg); err != nil {
					fmt.Println(err)
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
	}
	return commands
}


func cleanInput(text string) []string {
	cleanSlice := []string{}
	cleanSlice = strings.Fields(strings.ToLower(text))
	return  cleanSlice
}
