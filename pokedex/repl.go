package main
import "strings"
import "os"
import "bufio"
import "fmt"

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		result := cleanInput(scanner.Text())
		fmt.Printf("Your command was: %v \n", result[0])
	}

}

func cleanInput(text string) []string {
	cleanSlice := []string{}
	cleanSlice = strings.Fields(strings.ToLower(text))
	return  cleanSlice
}
