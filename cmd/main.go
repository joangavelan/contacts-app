package main

import (
	"log"
	"net/http"

	api "github.com/joangavelan/contacts-app/handlers/api"
	pages "github.com/joangavelan/contacts-app/handlers/pages"
	"github.com/joangavelan/contacts-app/internal/database"
)

func main() {
	// Create a new mux
	mux := http.NewServeMux()

	// Initialize database connection
	db, err := database.InitDB("contacts.db")
	if err != nil {
    log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Serve static files
	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	// group - pages
	mux.HandleFunc("/{$}", pages.Home)
	mux.HandleFunc("GET /auth/sign-in", pages.SignIn)
	mux.HandleFunc("GET /auth/sign-up", pages.SignUp)
	// group - api routes
	mux.HandleFunc("POST /api/register", api.RegisterUser)

	// Initialize server
	log.Fatal(http.ListenAndServe(":3000", mux))
}