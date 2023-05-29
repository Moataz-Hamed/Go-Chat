package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	con, err := sql.Open("postgres", "host=localhost password=password port=5432 user=postgres dbname=go-chat sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &Database{db: con}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
