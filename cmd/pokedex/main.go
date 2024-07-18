package main

import (
	"time"

	"github.com/meraiku/pokedex/internal/pokeapi"
	"github.com/meraiku/pokedex/internal/world"
)

type config struct {
	player                  world.Player
	pokeAPIClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeAPIClient: *pokeapi.NewClient(time.Hour),
		player:        *StartMsg(),
	}

	StartREPL(&cfg)
}
