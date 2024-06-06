package auth

import (
	"net/mail"
)

const (
	MinUsernameLength = 6
	MaxUsernameLength = 30
	MinPasswordLength = 6
	MaxPasswordLength = 50
	MaxEmailLength    = 100
)

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil && len(email) <= MaxEmailLength
}

func IsValidUsername(username string) bool {
	return len(username) >= MinUsernameLength && len(username) <= MaxUsernameLength
}

func IsValidPassword(password string) bool {
	return len(password) >= MinPasswordLength && len(password) <= MaxPasswordLength
}
