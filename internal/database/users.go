package database

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	Id        uuid.UUID
	Name      string
	Balance   float64
	CreatedAt time.Time
	Pokedex   Pokedex
}

type Pokedex struct {
	Pokedex map[string]Pokemons
}
