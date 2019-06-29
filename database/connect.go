package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type DBConnection struct {
	db    *sql.DB
	table string
}

func (conn *DBConnection) Close() {
	defer conn.Close()
}

var dbCache *sql.DB

func New(table string) (*DBConnection, error) {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &DBConnection{db: db, table: table}, nil
}

func (conn *DBConnection) CreateCustomerTable() error {
	queryCreateTable := `
 		CREATE TABLE IF NOT EXISTS ` + conn.table + `(
 			id SERIAL PRIMARY KEY,
 			name TEXT,
 			email TEXT,
 			status TEXT
 		);
 	`
	_, err := conn.db.Exec(queryCreateTable)
	if err != nil {
		return err
	}
	return nil
}
