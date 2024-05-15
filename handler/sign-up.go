package handler

import (
	"html/template"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"web/templates/base.html",
		"web/templates/auth/layout.html",
		"web/templates/commons/header.html",
		"web/templates/commons/footer.html",
		"web/templates/auth/sign-up.html",
	))
	
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
