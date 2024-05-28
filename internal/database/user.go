package database

import "fmt"

// CreateUser inserts a new user into the database and returns the ID of the newly inserted user
func CreateUser(username, email, hashedPassword string) (int64, error) {
	result, err := DB.Exec(InsertUser, username, email, hashedPassword)
	if err != nil {
		return 0, fmt.Errorf("failed to insert user %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return id, nil
}
