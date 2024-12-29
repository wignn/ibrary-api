package repository

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/wignn/library-api/internal/model"
)



func CreateGenre(db *DB, genre *model.Genre) error {
	existingGenre, err := GetGenreByName(db, genre.Name)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if existingGenre != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO genres (name) VALUES ($1)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(genre.Name)
	return err
}

func GetGenres(db *DB) ([]model.Genre, error) {
	stmt, err := db.Prepare("SELECT id, name, created_at FROM genres")
	if err != nil {
		return nil, fmt.Errorf("error preparing statement: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var genres []model.Genre
	for rows.Next() {
		var genre model.Genre
		err := rows.Scan(&genre.ID, &genre.Name, &genre.CreatedAt)
		if err != nil {
			log.Printf("error scanning genre: %v", err)
			return nil, fmt.Errorf("error scanning genre: %v", err)
		}
		genres = append(genres, genre)
	}
	return genres, nil
}

func GetGenreById(db *DB, id int) (*model.Genre, error) {
	var genre model.Genre

	stmt, err := db.Prepare("SELECT id, name, created_at FROM genres WHERE id = $1")
	if err != nil {
		return nil, fmt.Errorf("error preparing statement: %v", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&genre.ID, &genre.Name, &genre.CreatedAt)
	return &genre, err
}


func UpdateGenre(db *DB, id int, genre *model.Genre) error {
	stmt, err := db.Prepare("UPDATE genres SET name = $1 WHERE id = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(genre.Name, id)

	return err
}

func GetGenreByName(db *DB, name string) (*model.Genre, error) {
	var genre model.Genre
	stmt, err := db.Prepare("SELECT id, name, created_at FROM genres WHERE name = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(name).Scan(&genre.ID, &genre.Name, &genre.CreatedAt)

	return &genre, err
}


func DeleteGenre(db *DB, id int) error {
	stmt, err := db.Prepare("DELETE FROM genres WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)	
	return err
}

func AddGenreToBook(db *DB, id *model.BookGenre) error {
	stmt, err := db.Prepare("INSERT INTO book_genre (book_id, genre_id) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id.BookID, id.GenreID)
	return err
}

func GetBookGenres(db *DB, id int) ([]model.Genre, error) {
	stmt, err := db.Prepare(`SELECT genres.id, genres.name, genres.created_at FROM genres
	JOIN book_genre ON genres.id = book_genre.genre_id
	WHERE book_genre.book_id = $1 OR genres.id = $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []model.Genre
	for rows.Next() {
		var genre model.Genre
		if err := rows.Scan(&genre.ID, &genre.Name, &genre.CreatedAt); err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}
	return genres, err
}

// func GetBooksByGenre(db *DB, genreID int) ([]model.Book, error) {
// 	stmt, err := db.Prepare(`SELECT books.id, books.title, books.author, books.published_at FROM books
// 	JOIN book_genre ON books.id = book_genre.book_id
// 	WHERE book_genre.genre_id = $1`)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()

// 	rows, err := stmt.Query(genreID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var books []model.Book
// 	for rows.Next() {
// 		var book model.Book
// 		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.); err != nil {
// 			return nil, err
// 		}
// 		books = append(books, book)
// 	}
// 	return books, err
// }