package main

import (
	"time"

	"github.com/jhonipereira/pokecli/internal/api"
)

type config struct {
	apiClient       api.Client
	nextLocationURL *string
	prevLocationURL *string
	caughtPokemon   map[string]api.Pokemon
}

func main() {
	cfg := config{
		apiClient:     api.NewClient(time.Hour),
		caughtPokemon: make(map[string]api.Pokemon),
	}

	repl(&cfg)
}
