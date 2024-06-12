package handlers

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layouts/base.html",
		"web/templates/pages/home.html",
		"web/templates/commons/header.html",
		"web/templates/commons/footer.html",
	))

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
