package auth

import (
	"context"

	"github.com/joangavelan/contacts-app/internal/models"
)

// GetUser retrieves the UserContext from the provided context.
// It returns the UserContext and a boolean indicating whether the UserContext was found.
// This function is useful for accessing the user information stored in the request context by the auth middleware.
func GetUser(ctx context.Context) (*models.UserContext, bool) {
	userCtx, ok := ctx.Value(userContextKey).(*models.UserContext)
	return userCtx, ok
}
