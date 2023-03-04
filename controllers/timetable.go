package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"web-app/models"
)

// TimetableHandler renders a timetable page for a given user ID
func TimetableHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Get variables from URL path (requires gorilla/mux package)
	userID := vars["userID"]

	timetable, err := models.GetTimetableByUserID(userID) // Get timetable data from database (defined in models/timetable.go)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("views/layout.html", "views/timetable.html")) // Parse template files

	data := struct { // Define data to pass to template
		Title     string
		Timetable models.Timetable
	}{
		Title:     "Your Timetable",
		Timetable: timetable,
	}

	tmpl.ExecuteTemplate(w, "layout", data) // Execute template with data
}

// CreateTimetableHandler handles requests to create a new timetable from user input data (JSON format)
func CreateTimetableHandler(w http.ResponseWriter, r *http.Request) {
	var timetable models.Timetable // Define a timetable struct to hold user input data (defined in models/timetable.go)

	err := json.NewDecoder(r.Body).Decode(&timetable) // Decode JSON data from request body into timetable struct 
	if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return 
    }

	err = timetable.Create() // Save timetable data into database (defined in models/timetable.go)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return 
    }

	w.WriteHeader(http.StatusCreated) // Send back status code 201 (Created)
}
