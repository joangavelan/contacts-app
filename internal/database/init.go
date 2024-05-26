package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the database connection and assigns it to the global DB variable
func InitDB(dbName string) (*sql.DB, error) {
	var err error

	// Open DB connection
	dbPath := fmt.Sprintf("internal/database/%s", dbName)
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	return DB, nil
}