package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

// initialize the database
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		panic("Could not connect to database")
	}

	// how many open connection to have simultaneously to the database
	DB.SetMaxOpenConns(10)
	// how many connection to keep open if no one is using these connections at the moment
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	   id INTEGER PRIMARY KEY AUTOINCREMENT,
	   email TEXT NOT NULL UNIQUE,
	   password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	   id INTEGER PRIMARY KEY AUTOINCREMENT,
	   name TEXT NOT NULL,
	   description TEXT NOT NULL,
	   location TEXT NOT NULL,
	   dateTime DATETIME NOT NULL,
	   user_id INTEGER,
	   FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
	   id INTEGER PRIMARY KEY AUTOINCREMENT,
	   event_id INTEGER,
	   user_id INTEGER,
	   FOREIGN KEY(event_id) REFERENCES events(id),
	   FOREIGN KEY(user_id) REFERENCES users(id)
	)   
	`
	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not create registrations table")
	}

}
