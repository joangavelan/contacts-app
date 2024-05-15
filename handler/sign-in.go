package handler

import (
	"html/template"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"web/templates/base.html",
		"web/templates/auth/layout.html",
		"web/templates/commons/header.html",
		"web/templates/commons/footer.html",
		"web/templates/auth/sign-in.html",
	))
	
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
