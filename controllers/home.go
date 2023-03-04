package controllers

import (
	"html/template"
	"net/http"
)

// HomeHandler renders the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the template file
	tmpl := template.Must(template.ParseFiles("views/layout.html", "views/home.html"))

	// Execute the template with some data
	data := struct {
		Title string
	}{
		Title: "Welcome to Timetable Creator",
	}
	tmpl.ExecuteTemplate(w, "layout", data)
}
