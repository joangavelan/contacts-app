package handlers

import (
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layouts/base.html",
		"web/templates/layouts/auth.html",
		"web/templates/commons/header.html",
		"web/templates/commons/footer.html",
		"web/templates/pages/login/login.html",
		"web/templates/pages/login/form.html",
	))

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
