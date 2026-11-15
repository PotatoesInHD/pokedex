package main
import "strings"
import "os"
import "bufio"
import "fmt"


type cliCommand struct {
	name string
	description string
	callback func() error
}


func startRepl() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		result := cleanInput(scanner.Text())
		if command, ok := commands[result[0]]; ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
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
