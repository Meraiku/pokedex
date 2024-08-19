package database

import "database/sql"

type DB struct {
	db *sql.DB
}

func NewDB() DB {
	return DB{
		db: &sql.DB{},
	}
}
