package main

import (
	"time"

	"github.com/meraiku/pokedex/internal/pokeapi"
)

type config struct {
	pokeAPIClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeAPIClient: *pokeapi.NewClient(time.Hour),
	}

	StartREPL(&cfg)
}
