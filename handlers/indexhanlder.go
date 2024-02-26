package handlers

import (
	"html/template"
	"net/http"
)

// TO BE REMOVED
// Manages the index parsing and writing
// To remove in future implementations to show index.html as base static route
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./static/login.html")
	t.Execute(w, nil)
}
