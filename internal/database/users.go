package database

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/meraiku/pokedex/cmd/pokedex/structs"
	"github.com/uptrace/bun"
)

type Users struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	Id        uuid.UUID       `bun:"id,pk"`
	Name      string          `bun:"name,notnull,type:text"`
	Team      string          `bun:"team,notnull,type:text"`
	Balance   float64         `bun:"balance,notnull,default:100.0"`
	CreatedAt time.Time       `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	Pokedex   structs.Pokedex `bun:"-"`
}

func (db *DB) CreateUser(name, team string) error {
	user := &Users{
		Id:        uuid.New(),
		Name:      name,
		Team:      team,
		Pokedex:   structs.Pokedex{Pokedex: make(map[string]structs.Pokemon)},
		Balance:   100.0,
		CreatedAt: time.Now().UTC(),
	}

	_, err := db.db.NewInsert().Model(user).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetUserInfo() *Users {
	return &Users{}
}

func (db *DB) GetUsers() error {
	users := []Users{}

	err := db.db.NewSelect().Model(&users).Limit(10).Scan(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(users)
	return nil
}
