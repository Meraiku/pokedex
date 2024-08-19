package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/meraiku/pokedex/cmd/pokedex/structs"
)

type Users struct {
	Id        uuid.UUID
	Name      string
	Team      string
	Balance   float64
	CreatedAt time.Time
	Pokedex   structs.Pokedex
}

func (db *DB) GetUserInfo() *Users {
	return &Users{}
}
