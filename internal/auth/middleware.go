package auth

import (
	"context"
	"log"
	"net/http"

	"github.com/joangavelan/contacts-app/internal/models"
)

type contextKey string

const userContextKey = contextKey("user")

func Middleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve token from cookie
		cookie, err := r.Cookie("token")
		if err != nil {
			log.Printf("No token provided: %v", err)
			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}

		token := cookie.Value

		// Validate the token
		claims, err := ValidateJWT(token)
		if err != nil {
			log.Printf("Invalid token: %v", err)
			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}

		// Create a UserContext object from claims
		userCtx := &models.UserContext{
			Id:       claims.Sub,
			Username: claims.Username,
		}

		// Attach user context to request context
		ctx := context.WithValue(r.Context(), userContextKey, userCtx)

		// Proceed to the next handler with the new context
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
