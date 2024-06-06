package auth

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := "mysecretpassword"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	if hashedPassword == "" {
		t.Fatalf("expected a hashed password, but got an empty string")
	}

	// Ensure the hashed password is not equal to the plain password
	if hashedPassword == password {
		t.Fatalf("expected hashed password to be different from the plain password")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "mysecretpassword"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	// Test with the correct password
	err = CheckPasswordHash(password, hashedPassword)
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	// Test with an incorrect password
	incorrectPassword := "wrongpassword"
	err = CheckPasswordHash(incorrectPassword, hashedPassword)
	if err == nil {
		t.Fatalf("expected an error, but got none")
	}

	// Ensure the error is because of bcrypt comparison
	if err != bcrypt.ErrMismatchedHashAndPassword {
		t.Fatalf("expected bcrypt.ErrMismatchedHashAndPassword, but got %v", err)
	}
}
