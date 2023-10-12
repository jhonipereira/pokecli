package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func catch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.apiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("Failed to catch %s \n", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemonName)

	return nil
}
