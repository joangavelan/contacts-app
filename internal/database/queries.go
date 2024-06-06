package database

const (
	insertUserQuery = `
		INSERT INTO users (username, email, password)
		VALUES (?, ?, ?)
	`

	getUserQuery = `
		SELECT id, username, email, password FROM users WHERE email = ? LIMIT 1
	`

	emailExistsQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)
	`
)
