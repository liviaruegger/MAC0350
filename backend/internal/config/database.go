package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func SetupDatabase() *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

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
	);`

	activitiesTable := `
	CREATE TABLE IF NOT EXISTS activities (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL REFERENCES users(id),
		start TIMESTAMP NOT NULL,
		duration BIGINT NOT NULL,           -- in seconds
		distance FLOAT NOT NULL,
		laps INTEGER NOT NULL,
		pool_size FLOAT NOT NULL,
		location_type TEXT NOT NULL
	);`

	if _, err := db.Exec(userTable); err != nil {
		log.Fatal("Error creating users table:", err)
	}
	if _, err := db.Exec(activitiesTable); err != nil {
		log.Fatal("Error creating activities table:", err)
	}
}
