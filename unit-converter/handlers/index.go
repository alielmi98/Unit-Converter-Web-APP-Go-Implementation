package handlers

import (
	"html/template"
	"net/http"
)

// IndexHandler handles the index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}
