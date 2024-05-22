package handlers

import (
	"html/template"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layouts/base.html",
		"web/templates/layouts/auth.html",
		"web/templates/commons/header.html",
		"web/templates/commons/footer.html",
		"web/templates/pages/sign_up/sign_up.html",
		"web/templates/pages/sign_up/form.html",
	))
	
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
