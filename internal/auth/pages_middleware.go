package auth

import (
	"log"
	"net/http"
)

const (
	authCookieName = "token"
	redirectURL    = "/contacts"
)

func AuthPagesMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie(authCookieName)
		if err == nil {
			log.Println("Authenticated user trying to access auth page. Redirecting to main application page.")
			http.Redirect(w, r, redirectURL, http.StatusSeeOther)
			return
		} else if err != http.ErrNoCookie {
			// Log any error other than ErrNoCookie
			log.Printf("Error retrieving cookie: %v", err)
		}

		next.ServeHTTP(w, r)
	}
}
