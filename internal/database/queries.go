package database

const (
	InsertUser = `
		INSERT INTO users (username, email, password)
		VALUES (?, ?, ?)
	`

	emailExists = `
		SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)
	`
)