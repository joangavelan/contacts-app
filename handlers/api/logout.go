package handlers

import (
	"net/http"
	"time"
)

// Logout handles the logout process.
func Logout(w http.ResponseWriter, r *http.Request) {
	// Create an expired cookie to clear the JWT cookie
	expiredCookie := http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	}

	// Set the expired cookie in the response
	http.SetCookie(w, &expiredCookie)

	// Redirect to login page
	w.Header().Set("HX-Redirect", "/auth/login")
	w.WriteHeader(http.StatusSeeOther)
}
