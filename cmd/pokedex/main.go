package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}

	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cfg := config{
		pokeAPIClient: *pokeapi.NewClient(time.Hour),
		player:        *StartMsg(),
	}
	StartREPL(&cfg)
}
