package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/meraiku/pokedex/internal/database"
)

type config struct {
	db                      *database.DB
	pokeAPIClient           Client
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
		pokeAPIClient: *NewClient(time.Hour),
		db:            database.NewDB(db),
	}
	StartREPL(&cfg)
}
