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

// Login handles the user login process by validating inputs,
// checking credentials, generating a JWT, and setting the token in a cookie.
func Login(w http.ResponseWriter, r *http.Request) {
	// Parse and Validate Form Data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	LoginForm := models.LoginForm{}
	LoginForm.Values.Email = strings.TrimSpace(r.FormValue("email"))
	LoginForm.Values.Password = strings.TrimSpace(r.FormValue("password"))

	if !auth.IsValidEmail(LoginForm.Values.Email) {
		LoginForm.Errors.Email = "Invalid email address"
	}

	if !auth.IsValidPassword(LoginForm.Values.Password) {
		LoginForm.Errors.Password = fmt.Sprintf("Password must be between %d and %d characters long", auth.MinPasswordLength, auth.MaxPasswordLength)
	}

	// Render form with errors and submitted values if validation fails.
	if LoginForm.HasErrors() {
		tmpl := template.Must(template.ParseFiles("web/templates/pages/login/form.html"))
		if err := tmpl.Execute(w, LoginForm); err != nil {
			http.Error(w, "Unable to render template", http.StatusInternalServerError)
		}
		return
	}

	// Retrieve user from database.
	user, err := database.GetUserByEmail(database.DB, LoginForm.Values.Email)
	if err != nil {
		log.Printf("Error retrieving user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if user == nil {
		if err := toast.Error("User not found").WriteToHeader(w); err != nil {
			log.Printf("Error writing toast event: %v", err)
		}
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Verify password.
	if err := auth.CheckPasswordHash(LoginForm.Values.Password, user.Password); err != nil {
		if err := toast.Error("Incorrect Password").WriteToHeader(w); err != nil {
			log.Printf("Error writing toast event: %v", err)
		}
		http.Error(w, "Incorrect Password", http.StatusUnauthorized)
		return
	}

	// Generate JWT.
	tokenString, err := auth.GenerateJWT(user.Id, user.Username, user.Email)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		http.Error(w, "Error generating JWT", http.StatusInternalServerError)
		return
	}

	// Set JWT in a cookie.
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		Expires:  time.Now().Add(config.CookieExpiration),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	})

	// Redirect to contacts page.
	w.Header().Set("HX-Redirect", "/contacts")
	w.WriteHeader(302)
}
