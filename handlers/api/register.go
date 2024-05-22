package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strings"
)

type SignUpFormFields struct { 
	Username string
	Email string
	Password string
}

type SignUpForm struct {
	Values SignUpFormFields
	Errors SignUpFormFields
}

// Constants for validation
const (
	MinUsernameLength = 6
	MaxUsernameLength = 30
	MinPasswordLength = 6
	MaxPasswordLength = 50
	MaxEmailLength    = 100
)

// Helper function to validate email using regex
func isValidEmail(email string) bool {
	// Simple regex for email validation
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	signUpForm := SignUpForm{}

	// Retrieve and trim form values
	signUpForm.Values.Username = strings.TrimSpace(r.FormValue("username"))
	signUpForm.Values.Email = strings.TrimSpace(r.FormValue("email"))
	signUpForm.Values.Password = strings.TrimSpace(r.FormValue("password"))

	// Validate username
	if len(signUpForm.Values.Username) < MinUsernameLength || len(signUpForm.Values.Username) > MaxUsernameLength {
		signUpForm.Errors.Username = fmt.Sprintf("Username must be between %d and %d characters long", MinUsernameLength, MaxUsernameLength)
	}

	// Validate email
	if len(signUpForm.Values.Email) > MaxEmailLength || !isValidEmail(signUpForm.Values.Email) {
		signUpForm.Errors.Email = "Invalid email address"
	}

	// Validate password
	if len(signUpForm.Values.Password) < MinPasswordLength || len(signUpForm.Values.Password) > MaxPasswordLength {
		signUpForm.Errors.Password = fmt.Sprintf("Password must be between %d and %d characters long", MinPasswordLength, MaxPasswordLength)
	}

	// Check if any errors
	if signUpForm.Errors.Username != "" || signUpForm.Errors.Email != "" || signUpForm.Errors.Password != "" {
		// Render the form with errors and preserved input values
		tmpl := template.Must(template.ParseFiles("web/templates/pages/sign_up/form.html"))
		if err := tmpl.Execute(w, signUpForm); err != nil {
			http.Error(w, "Unable to render template", http.StatusInternalServerError)
		}
		return
	}

	// Respond with a success message (replace this with a redirect)
	fmt.Fprintf(w, `<div>Registration successful! Redirecting...</div>`)
}
