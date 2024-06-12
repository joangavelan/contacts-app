package handlers

import (
	"html/template"
	"net/http"

	"github.com/joangavelan/contacts-app/internal/auth"
)

func Contacts(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layouts/base.html",
		"web/templates/pages/contacts/contacts.html",
	))

	user, ok := auth.GetUser(r.Context())
	if !ok {
		http.Error(w, "Could not retrieve user information", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
