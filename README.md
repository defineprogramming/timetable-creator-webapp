# Timetable Creator Web App

This is a simple web app built with Go that allows users to create and view their revision timetables.

## Features

- Users can sign up with their name, email and password.
- Users can log in with their email and password.
- Users can view their profile with their name and email.
- Users can create a timetable for a subject by entering the subject name, start day and end day of the revision period.
- Users can view their timetable with the subject name, start day, end day and a list of topics to revise for each date.
- Users can log out from their account.

## Installation

To run this web app locally, you need to have Go installed on your machine. You also need to have a MongoDB database running on your localhost:27017 port.

1. Clone this repository to your local machine using `git clone https://github.com/defineprogramming/timetable-creator-webapp/`.
2. Navigate to the web-app folder using `cd timetable-creator-webapp`.
3. Install the required dependencies using `github.com/gorilla/mux.`.
4. Run the main.go file using `go run main.go`.
5. Open your browser and go to http://localhost:8080 to see the web app in action.

## File Structure

The web app consists of the following files and folders:

- main.go: This file contains the main function that initializes the database connection, defines the routes and handlers for each route, and starts the server on port 8080.
- controller.go: This file contains the controller functions that handle each route request, perform database operations, render HTML templates or redirect to other routes as needed.
- model.go: This file contains the model structs that represent the data structures for users and timetables in MongoDB collections.
- views: This folder contains HTML template files that are rendered by the controller functions. Each template file can use Go's template syntax to inject data from the controller into the HTML.
- static: This folder contains static files such as CSS, JavaScript and images that are served by the server.

## License

This project is licensed under the MIT License - see [LICENSE](LICENSE) file for details.
