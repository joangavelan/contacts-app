package database

import (
	"database/sql"
	"fmt"

	"github.com/joangavelan/contacts-app/internal/models"
)

// CreateUser inserts a new user into the database and returns the ID of the newly inserted user.
func CreateUser(db *sql.DB, username, email, hashedPassword string) (int64, error) {
	result, err := db.Exec(insertUserQuery, username, email, hashedPassword)
	if err != nil {
		return 0, fmt.Errorf("failed to insert user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return id, nil
}

// GetUserByEmail retrieves a user from the database by their email address.
func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
	row := db.QueryRow(getUserQuery, email)

	var user models.User
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, fmt.Errorf("failed to query user: %w", err)
	}

	return &user, nil
}

// EmailExists checks if the provided email already exists in the users table.
func EmailExists(db *sql.DB, email string) (bool, error) {
	var exists bool

	err := db.QueryRow(emailExistsQuery, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
