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
		log.Printf("Error creating tables: %v\n", err)
		return nil, err
	}
	
	log.Println("Tables created successfully.")
	return &DB{db}, nil
}


func createTable(db *sql.DB) error {
	migrations := []func(*sql.DB) error{
		createUsersTable,
		createBooksTable,
		createGenresTable,
		createBookGenreTable,
		createBookmarksTable,
		createReviewsTable,
		createUserReviewsTable,
	}

	for i, migration := range migrations {
		log.Printf("Applying migration %d...\n", i+1)
		if err := migration(db); err != nil {
			log.Printf("Migration %d failed: %v\n", i+1, err)
			return err
		}
	}

	log.Println("All migrations applied successfully.")
	return nil
}
func createUsersTable(db *sql.DB) error {
_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY NOT NULL,
	username VARCHAR(255) NOT NULL UNIQUE,
	email VARCHAR(255) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	is_admin BOOLEAN DEFAULT FALSE,
	profile_picture TEXT,
	is_active BOOLEAN DEFAULT TRUE,
	token TEXT
)`)
	return err
}

func createBooksTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY NOT NULL,
		title VARCHAR(255) NOT NULL,
		author VARCHAR(255) NOT NULL,
		published_date DATE NOT NULL,
		description TEXT,
		cover VARCHAR(255),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	return err
}
func createGenresTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS genres (
		id SERIAL PRIMARY KEY NOT NULL,
		name VARCHAR(255) NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	return err
}

func createBookGenreTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS book_genre (
		book_id INT NOT NULL,
		genre_id INT NOT NULL,
		PRIMARY KEY (book_id, genre_id),
		FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
		FOREIGN KEY (genre_id) REFERENCES genres(id) ON DELETE CASCADE
	)`)
	return err
}

func createBookmarksTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS bookmarks (
		id SERIAL PRIMARY KEY NOT NULL,
		book_id INT NOT NULL,
		user_id INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	)`)
	return err
}

func createReviewsTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS reviews (
		id SERIAL PRIMARY KEY NOT NULL,
		content TEXT NOT NULL,
		is_spoiler BOOLEAN DEFAULT FALSE,
		book_id INT NOT NULL,
		user_id INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	)`)
	return err
}

func createUserReviewsTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS user_reviews (
		id SERIAL PRIMARY KEY NOT NULL,
		content TEXT NOT NULL,
		is_spoiler BOOLEAN DEFAULT FALSE,
		user_id INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	)`)
	return err
}
