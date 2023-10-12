package main

import "fmt"

func helpCmd(cfg *config, args ...string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, help := range getCommands() {
		fmt.Println(help.name + ": " + help.description + "\n")
	}
	return nil
}
