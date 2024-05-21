package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Define SQL statements for table creation
const (
	createUsersTableSQLQuery = `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		email TEXT NOT NULL
	)`

	createContactsTableSQLQuery = `CREATE TABLE IF NOT EXISTS contacts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		firstName TEXT NOT NULL,
		lastName TEXT NOT NULL,
		email TEXT NOT NULL,
		phoneNumber TEXT NOT NULL,
		userId INTEGER NOT NULL,
			FOREIGN KEY (userId) REFERENCES users(id)
	)`
)

func SetupDB(dbName string) (*sql.DB, error) {
	// Open DB connection
	db, err := sql.Open("sqlite3", "database/" + dbName)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	// Create users table
	if _, err := db.Exec(createUsersTableSQLQuery); err != nil {
		return nil, fmt.Errorf("error creating users table: %v", err)
	}

	// Create contacts table
	if _, err := db.Exec(createContactsTableSQLQuery); err != nil {
		return nil, fmt.Errorf("error creating contacts table: %v", err)
	}

	return db, nil
}