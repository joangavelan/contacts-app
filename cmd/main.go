package main

import (
	"log"
	"net/http"

	"github.com/joangavelan/contacts-app/database"
	api "github.com/joangavelan/contacts-app/handlers/api"
	pages "github.com/joangavelan/contacts-app/handlers/pages"
)

func main() {
	// Create a new mux
	mux := http.NewServeMux()

	// Setup database
	db, err := database.SetupDB("contacts.db")
	if err != nil {
		log.Fatalf("Error setting up database: %v", err)
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