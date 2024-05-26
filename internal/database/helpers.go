package database

// EmailExists checks if the provided email already exists in the users table
func EmailExists(email string) (bool, error) {
	var exists bool

	err := DB.QueryRow(emailExists, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
