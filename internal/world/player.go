package world

import (
	"time"

	"github.com/meraiku/pokedex/internal/pokeapi"
)

type Player struct {
	Name      string
	Age       int
	Balance   float64
	Pokedex   Pokedex
	createdAt time.Time
}

type Pokedex struct {
	Pokedex map[string]pokeapi.Pokemon
}

func NewPlayer(name string, age int) *Player {
	return &Player{
		Name:      name,
		Age:       age,
		Balance:   100.0,
		Pokedex:   *NewPokedex(),
		createdAt: time.Now(),
	}
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		Pokedex: make(map[string]pokeapi.Pokemon),
	}
}
