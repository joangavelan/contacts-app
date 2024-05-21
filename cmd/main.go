package main

import (
	"log"
	"net/http"

	"github.com/joangavelan/contacts-app/database"
	"github.com/joangavelan/contacts-app/handlers"
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
	mux.HandleFunc("/{$}", handlers.Home)
	mux.HandleFunc("GET /auth/sign-in", handlers.SignIn)
	mux.HandleFunc("GET /auth/sign-up", handlers.SignUp)

	// Initialize server
	log.Fatal(http.ListenAndServe(":3000", mux))
}