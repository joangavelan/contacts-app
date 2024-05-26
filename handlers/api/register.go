package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/joangavelan/contacts-app/internal/database"
	"github.com/joangavelan/contacts-app/internal/models"

	"golang.org/x/crypto/bcrypt"
)

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

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	signUpForm := models.SignUpForm{}

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

	// Store user information securely in the database
	// 1. Verify if the email already exists
	exists, err := database.EmailExists(signUpForm.Values.Email)
	if err != nil {
    log.Printf("Error checking email existence: %v", err)
		http.Error(w, "Error checking email existence", http.StatusInternalServerError)
		return
	}

	if exists {
		// Respond with a 409 Conflict status if the email is already registered
		http.Error(w, "Email address already registered", http.StatusConflict)
		return
	}
	
	// 2. Hash the password for safe storage
	hashedPassword, err := hashPassword(signUpForm.Values.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	
	// 3. Insert user data into the database
	_, err = database.DB.Exec(database.InsertUser, signUpForm.Values.Username, signUpForm.Values.Email, hashedPassword)
	if err != nil {
		// Log the specific error for debugging purposes
    log.Printf("Error inserting user data into database: %v", err)
		http.Error(w, "Error storing user data", http.StatusInternalServerError)
    return
	}

	// Respond with a success message (replace this with a redirect)
	fmt.Fprintf(w, `<div>Registration successful! Redirecting...</div>`)
}
