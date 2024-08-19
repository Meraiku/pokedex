package main

import (
	"database/sql"
	"fmt"
	"os"

	"errors"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func connectDB() (*bun.DB, error) {
	DB_URL := os.Getenv("DATABASE_URL")
	if DB_URL == "" {
		return nil, errors.New("DATABASE_URL environment varible is not set")
	}

	fmt.Println("Statring connection with Postgres")

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(DB_URL)))

	db := bun.NewDB(sqldb, pgdialect.New())

	if err := db.Ping(); err != nil {
		fmt.Println("DB Ping failed")
		return nil, err
	}

	fmt.Println("DB Connection online!")
	return db, nil
}
