package controllers

import (
    "encoding/json"
    "html/template"
    "net/http"
    "web-app/models"
)

// LoginHandler handles requests to log in a user with email and password
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var user models.User // Define a user struct to hold user input data (defined in models/user.go)

    err := json.NewDecoder(r.Body).Decode(&user) // Decode JSON data from request body into user struct
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    token, err := user.Login() // Validate user credentials and generate a JWT token (defined in models/user.go)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    w.Header().Set("Content-Type", "application/json") // Set response header to JSON format
    w.WriteHeader(http.StatusOK) // Send back status code 200 (OK)
    
    json.NewEncoder(w).Encode(token) // Encode token as JSON and write to response body
}

// ProfileHandler renders a profile page for a logged-in user
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
    
    token := r.Header.Get("Authorization") // Get token from request header

    user, err := models.GetUserByToken(token) // Get user data from database using token (defined in models/user.go)
    if err != nil {
        http.Error(w, err.Error(), http.StatusForbidden)
        return
    }

	tmpl := template.Must(template.ParseFiles("views/layout.html", "views/user.html")) // Parse template files

	data := struct { // Define data to pass to template
		Title string
		User  models.User
	}{
		Title: "Your Profile",
		User:  user,
	}

	tmpl.ExecuteTemplate(w, "layout", data) // Execute template with data

}
