package main
//import "fmt"

func main() {
	commands := getCommands()
	cfg := &config{commands: commands}
	startRepl(cfg)
}
