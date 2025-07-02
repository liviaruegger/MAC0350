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
		id UUID PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		city TEXT NOT NULL,
		phone TEXT NOT NULL
	);`

	activitiesTable := `
	CREATE TABLE IF NOT EXISTS activities (
		id UUID PRIMARY KEY,
		user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		start TIMESTAMP NOT NULL,
		duration BIGINT NOT NULL,
		distance FLOAT NOT NULL,
		laps INTEGER NOT NULL,
		pool_size FLOAT NOT NULL,
		location_type TEXT NOT NULL CHECK (location_type IN ('pool', 'open_water')),
		notes TEXT DEFAULT ''
	);`

	intervalsTable := `
	CREATE TABLE IF NOT EXISTS intervals (
		id UUID PRIMARY KEY,
		activity_id UUID NOT NULL REFERENCES activities(id) ON DELETE CASCADE,
		duration BIGINT NOT NULL,
		distance FLOAT NOT NULL,
		type TEXT NOT NULL CHECK (
			type IN (
				'swim', 'rest', 'drill', 'kick', 'pull',
				'warmup', 'main_set', 'cooldown'
			)
		),
		stroke TEXT NOT NULL CHECK (
			stroke IN (
				'freestyle', 'backstroke', 'breaststroke',
				'butterfly', 'medley', 'unknown'
			)
		),
		notes TEXT DEFAULT ''
	);`

	if _, err := db.Exec(userTable); err != nil {
		log.Fatal("Error creating users table:", err)
	}
	if _, err := db.Exec(activitiesTable); err != nil {
		log.Fatal("Error creating activities table:", err)
	}
	if _, err := db.Exec(intervalsTable); err != nil {
		log.Fatal("Error creating intervals table:", err)
	}
}
