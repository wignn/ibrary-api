package repository

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func InitDb() (*DB, error) {

	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Printf("Error opening database connection: %v\n", err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		log.Printf("Error pinging database: %v\n", err)
		return nil, err
	}
	log.Println("Database connection successfully initialized.")
	err = createTable(db)
	if err != nil {
		log.Printf("Error creating table: %v\n", err)
		return nil, err
	}
	return &DB{db}, nil
}


func createTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY NOT NULL,
		username VARCHAR(255) NOT NULL UNIQUE,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		is_admin BOOLEAN DEFAULT FALSE,
		profile_picture TEXT,
		is_active BOOLEAN DEFAULT TRUE
	)`)
	return err
}
