package models

import (
	"database/sql"
)

// Timetable represents a revision timetable for a given subject and date range
type Timetable struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"user_id"`
	Subject  string `json:"subject"`
	StartDay string `json:"start_day"`
	EndDay   string `json:"end_day"`
}

// GetTimetableByUserID returns a timetable for a given user ID from the database 
func GetTimetableByUserID(userID int64) (Timetable, error) {
	var timetable Timetable

	db, err := sql.Open("sqlite3", "./web-app.db") // Open database connection (requires sqlite3 driver package)
	if err != nil {
		return timetable, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT id, subject, start_day, end_day FROM timetables WHERE user_id = ?", userID) // Query database for timetable by user ID
	
	err = row.Scan(&timetable.ID, &timetable.Subject, &timetable.StartDay,&timetable.EndDay) // Scan row into timetable struct 
	if err != nil {
		return timetable,err 
     }

	timetable.UserID = userID 

	return timetable,nil 
}

// Create saves a new timetable into the database 
func (t *Timetable) Create() error {

	db,err:=sql.Open("sqlite3","./web-app.db")// Open database connection (requires sqlite3 driver package)
	if err!=nil{
			returnerr 
     }
	deferdb.Close()

	result,err:=db.Exec("INSERT INTO timetables(user_id ,subject,start_day,end_day )VALUES(?,?,?,?)",t.UserID,t.Subject,t.StartDay,t.EndDay)// Insert new record into timetables table 
	iferr!=nil{
			returnerr 
     }

	t.ID,err=result.LastInsertId()// Get last inserted ID and assign it to t.ID 
	iferr!=nil{
			returnerr 
     }

	returnnil 

}
