package main

import (
	"log"
	"net/http"

	"github.com/joangavelan/contacts-app/database"
	"github.com/joangavelan/contacts-app/handler"
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
	mux.HandleFunc("/{$}", handler.Home)
	mux.HandleFunc("GET /auth/sign-in", handler.SignIn)
	mux.HandleFunc("GET /auth/sign-up", handler.SignUp)

	// Initialize server
	log.Fatal(http.ListenAndServe(":3000", mux))
}