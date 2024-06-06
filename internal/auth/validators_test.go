package auth

import (
	"strings"
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},
		{"user.name+tag+sorting@example.com", true},
		{"user@example", true},
		{"user@.com", false},
		{"user@com", true},
		{"user@example..com", false},
		{"", false},
		{"a@b.c", true},
		{"valid.email@example.com", true},
		{"email@subdomain.example.com", true},
		{"email@example.museum", true},
		{"email@example.co.jp", true},
		{"email@example.com", true},
		{"email@-example.com", true},
		{"plainaddress", false},
		{"@missinguser.com", false},
		{strings.Repeat("a", MaxEmailLength-12) + "@example.com", true},  // email with max length
		{strings.Repeat("a", MaxEmailLength-11) + "@example.com", false}, // 1 character too long
	}

	for _, test := range tests {
		result := IsValidEmail(test.email)
		if result != test.expected {
			t.Errorf("IsValidEmail(%q) = %v; want %v", test.email, result, test.expected)
		}
	}
}

func TestIsValidUsername(t *testing.T) {
	tests := []struct {
		username string
		expected bool
	}{
		{"short", false},                                  // 5 characters
		{"validusername", true},                           // 13 characters
		{strings.Repeat("a", MinUsernameLength), true},    // exactly 6 characters
		{strings.Repeat("a", MaxUsernameLength), true},    // exactly 30 characters
		{strings.Repeat("a", MaxUsernameLength+1), false}, // 31 characters
		{"", false},              // empty username
		{"valid_username", true}, // valid username with underscore
	}

	for _, test := range tests {
		result := IsValidUsername(test.username)
		if result != test.expected {
			t.Errorf("IsValidUsername(%q) = %v; want %v", test.username, result, test.expected)
		}
	}
}

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"short", false},                                  // 5 characters
		{"validpassword", true},                           // 12 characters
		{strings.Repeat("a", MinPasswordLength), true},    // exactly 6 characters
		{strings.Repeat("a", MaxPasswordLength), true},    // exactly 50 characters
		{strings.Repeat("a", MaxPasswordLength+1), false}, // 51 characters
		{"", false},                 // empty password
		{"valid_password123", true}, // valid password with special characters and numbers
	}

	for _, test := range tests {
		result := IsValidPassword(test.password)
		if result != test.expected {
			t.Errorf("IsValidPassword(%q) = %v; want %v", test.password, result, test.expected)
		}
	}
}
