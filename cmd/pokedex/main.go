package main

import (
	"github.com/meraiku/pokedex/internal/pokeapi"
)

type config struct {
	pokeAPIClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeAPIClient: *pokeapi.NewClient(),
	}

	StartREPL(&cfg)
}
