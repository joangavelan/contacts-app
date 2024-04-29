package handler

import (
	"net/http"
	"text/template"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"web/templates/base.html",
		"web/templates/auth/layout.html",
		"web/templates/auth/sign-up.html",
	))
	
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
