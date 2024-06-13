package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/joangavelan/contacts-app/config"
	"github.com/joangavelan/contacts-app/internal/auth"
	"github.com/joangavelan/contacts-app/internal/database"
	"github.com/joangavelan/contacts-app/internal/models"
	"github.com/joangavelan/contacts-app/pkg/toast"
)

// Register handles user registration by validating inputs,
// storing user data, generating a JWT, and setting the token in a cookie.
func Register(w http.ResponseWriter, r *http.Request) {
	// Parse and Validate Form Inputs
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	registerForm := models.RegisterForm{}
	registerForm.Values.Username = strings.TrimSpace(r.FormValue("username"))
	registerForm.Values.Email = strings.TrimSpace(r.FormValue("email"))
	registerForm.Values.Password = strings.TrimSpace(r.FormValue("password"))

	if !auth.IsValidUsername(registerForm.Values.Username) {
		registerForm.Errors.Username = fmt.Sprintf("Username must be between %d and %d characters long", auth.MinUsernameLength, auth.MaxUsernameLength)
	}

	if !auth.IsValidEmail(registerForm.Values.Email) {
		registerForm.Errors.Email = "Invalid email address"
	}

	if !auth.IsValidPassword(registerForm.Values.Password) {
		registerForm.Errors.Password = fmt.Sprintf("Password must be between %d and %d characters long", auth.MinPasswordLength, auth.MaxPasswordLength)
	}

	// Render form with errors and submitted values if validation fails.
	if registerForm.HasErrors() {
		tmpl := template.Must(template.ParseFiles("web/templates/pages/register/form.html"))
		if err := tmpl.Execute(w, registerForm); err != nil {
			http.Error(w, "Unable to render template", http.StatusInternalServerError)
		}
		return
	}

	// Store user information securely.
	exists, err := database.EmailExists(database.DB, registerForm.Values.Email)
	if err != nil {
		log.Printf("Error checking email existence: %v", err)
		http.Error(w, "Error checking email existence", http.StatusInternalServerError)
		return
	}

	if exists {
		if err := toast.Error("Email address already registered").WriteToHeader(w); err != nil {
			log.Printf("Error writing toast event: %v", err)
		}
		http.Error(w, "Email address already registered", http.StatusConflict)
		return
	}

	hashedPassword, err := auth.HashPassword(registerForm.Values.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	userId, err := database.CreateUser(database.DB, registerForm.Values.Username, registerForm.Values.Email, hashedPassword)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// JWT creation and delivery.
	tokenString, err := auth.GenerateJWT(userId, registerForm.Values.Username, registerForm.Values.Email)
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
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	// Redirect to contacts page.
	w.Header().Set("HX-Redirect", "/contacts")
	w.WriteHeader(http.StatusSeeOther)
}
