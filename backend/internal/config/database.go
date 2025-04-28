package config

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func SetupDatabase() *sql.DB {
    dsn := "host=localhost port=5432 user=postgres password=postgres dbname=tracker sslmode=disable"
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal("Database connection error:", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatal("Database ping error:", err)
    }

    createTables(db)

    return db
}

func createTables(db *sql.DB) {
    userTable := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        email TEXT UNIQUE NOT NULL
    );
    `

    activitiesTable := `
    CREATE TABLE IF NOT EXISTS activities (
        id SERIAL PRIMARY KEY,
        distance TEXT NOT NULL,
        start INT NOT NULL,
		finish INT NOT NULL,
		size INT NOT NULL,
		laps INT NOT NULL,
        user_id INT NOT NULL REFERENCES users(id)
    );
    `

    if _, err := db.Exec(userTable); err != nil {
        log.Fatal("Error creating users table:", err)
    }
    if _, err := db.Exec(activitiesTable); err != nil {
        log.Fatal("Error creating activities table:", err)
    }
}