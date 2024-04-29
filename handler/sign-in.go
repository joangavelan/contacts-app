package handler

import (
	"net/http"
	"text/template"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"web/templates/base.html",
		"web/templates/auth/layout.html",
		"web/templates/auth/sign-in.html",
	))
	
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
