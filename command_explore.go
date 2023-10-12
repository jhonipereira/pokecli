package main

import (
	"errors"
	"fmt"
)

func explore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	locationAreaName := args[0]

	locationArea, err := cfg.apiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("Explore all the Pokemons in %s:\n", locationArea.Name)
	for _, poke := range locationArea.PokemonEncounters {
		fmt.Printf(" -  %s\n", poke.Pokemon.Name)
	}

	return nil
}
