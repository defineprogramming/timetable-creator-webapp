package main

import (
	"net/http"
	"web-app/controllers"
)

func main() {
	// Create a new router
	r := controllers.NewRouter()

	// Start listening for incoming requests
	http.ListenAndServe(":8080", r)
}
