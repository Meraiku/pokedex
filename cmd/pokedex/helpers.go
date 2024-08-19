package main

import (
	"database/sql"
	"fmt"
	"os"

	"errors"
)

func connectDB() (*sql.DB, error) {
	DB_URL := os.Getenv("DATABASE_URL")
	if DB_URL == "" {
		return nil, errors.New("DATABASE_URL environment varible is not set")
	}

	fmt.Println("Statring connection with Postgres")

	db, err := sql.Open("postgres", DB_URL)
	if err != nil {
		return nil, fmt.Errorf("error connecting database: %s", err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println("DB Ping failed")
		return nil, err
	}

	fmt.Println("DB Connection online!")
	return db, nil
}
