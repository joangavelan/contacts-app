package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/joangavelan/contacts-app/config"
	"github.com/joangavelan/contacts-app/internal/auth"
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
	// *******************************************
	// * STEP 1: PARSE AND VALIDATE FORM INPUTS *
	// *******************************************
	// Extract form data from the HTTP request and validate each field
	// Ensure that required fields are present and adhere to expected formats
	// Render the form with validation errors and input values if any errors are present
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	signUpForm := models.SignUpForm{}

	signUpForm.Values.Username = strings.TrimSpace(r.FormValue("username"))
	signUpForm.Values.Email = strings.TrimSpace(r.FormValue("email"))
	signUpForm.Values.Password = strings.TrimSpace(r.FormValue("password"))

	if len(signUpForm.Values.Username) < MinUsernameLength || len(signUpForm.Values.Username) > MaxUsernameLength {
		signUpForm.Errors.Username = fmt.Sprintf("Username must be between %d and %d characters long", MinUsernameLength, MaxUsernameLength)
	}

	if len(signUpForm.Values.Email) > MaxEmailLength || !isValidEmail(signUpForm.Values.Email) {
		signUpForm.Errors.Email = "Invalid email address"
	}

	if len(signUpForm.Values.Password) < MinPasswordLength || len(signUpForm.Values.Password) > MaxPasswordLength {
		signUpForm.Errors.Password = fmt.Sprintf("Password must be between %d and %d characters long", MinPasswordLength, MaxPasswordLength)
	}

	if signUpForm.Errors.Username != "" || signUpForm.Errors.Email != "" || signUpForm.Errors.Password != "" {
		tmpl := template.Must(template.ParseFiles("web/templates/pages/sign_up/form.html"))
		if err := tmpl.Execute(w, signUpForm); err != nil {
			http.Error(w, "Unable to render template", http.StatusInternalServerError)
		}
		return
	}

	// ************************************************************
	// * STEP 2: STORE USER INFORMATION SECURELY IN THE DATABASE *
	// ************************************************************
	// Ensure the email is unique to prevent duplicate accounts
	// Insert the validated user data into the database
	// Handle any database errors and provide appropriate feedback to the client
	exists, err := database.EmailExists(signUpForm.Values.Email)
	if err != nil {
		log.Printf("Error checking email existence: %v", err)
		http.Error(w, "Error checking email existence", http.StatusInternalServerError)
		return
	}

	if exists {
		http.Error(w, "Email address already registered", http.StatusConflict)
		return
	}

	hashedPassword, err := hashPassword(signUpForm.Values.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	userId, err := database.CreateUser(signUpForm.Values.Username, signUpForm.Values.Email, hashedPassword)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		http.Error(w, "Error creating new user", http.StatusInternalServerError)
		return
	}

	// ********************************************************************
	// * STEP 3: JWT CREATION AND DELIVERY AFTER SUCCESSFUL REGISTRATION *
	// ********************************************************************
	// Generate a JWT token for the newly registered user
	// Set the token as a secure HTTP-only cookie
	tokenString, err := auth.GenerateJWT(userId, signUpForm.Values.Email, signUpForm.Values.Username)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		http.Error(w, "Error generating JWT", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		Expires:  time.Now().Add(config.CookieExpiration),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	})

	// Respond with a success message (replace this with a redirect)
	fmt.Fprintf(w, `<div>Registration successful! Redirecting...</div>`)
}
