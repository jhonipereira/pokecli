package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    nil,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    nil,
		},
	}
}

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your command: ")
		cmd, _ := reader.ReadString('\n')
		fmt.Print("Your command was " + cmd)
		if cmd == "help" {
			getCommands()
		}
	}
}
