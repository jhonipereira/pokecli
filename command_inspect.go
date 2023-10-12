package main

import (
	"errors"
	"fmt"
)

func inspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name to be inspected")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("haven't caught the pokemon yet")
	}
	fmt.Printf("ID: %v\n", pokemon.ID)
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Print("STATS: \n", pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	for _, typ := range pokemon.Types {
		fmt.Printf(" Type %s\n", typ.Type.Name)
	}

	return nil
}
