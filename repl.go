package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCmd,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exitCmd,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Each time you call map, it displays the next 20 locations",
			callback:    mapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world. It's the 'reverse' direction of the map command.",
			callback:    mapBack,
		},
		"explore": {
			name:        "explore {location_name/id}",
			description: "Explore the specific area (id or name). To see the available areas, run the 'map' or 'mapb' command.",
			callback:    explore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Catch the specific Pokemon if your experience is enough against the level of the Pokemon.",
			callback:    catch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "View the information about a specific Pokemon.",
			callback:    inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View all Pokemons that you've caught.",
			callback:    pokedex,
		},
	}
}

func repl(cfg *config) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("POKEDEX > ")
		words, _ := reader.ReadString('\n')
		sanitizedWord := FilterWord(words)
		if len(sanitizedWord) == 0 {
			continue
		}
		cmd := sanitizedWord[0]
		args := []string{}
		if len(sanitizedWord) > 1 {
			args = sanitizedWord[1:]
		}
		commands := getCommands()
		if _, ok := commands[cmd]; ok {
			err := commands[cmd].callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("command not found: %v\n", cmd)
			continue
		}
	}
}

func FilterWord(str string) []string {
	low := strings.ToLower(str)
	words := strings.Fields(low)
	return words
}
