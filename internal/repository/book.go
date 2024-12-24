package repository

import (
	"github.com/wignn/library-api/internal/model"
	"log"
	"time"
)

func CreateBook(db *DB, book *model.Book) error {
	_, err := db.Exec(`INSERT INTO books (title, author, published_date, description, cover) VALUES ($1, $2, $3, $4, $5)`, book.Title, book.Author, book.PublisedDate, book.Description, book.Cover)
	return err
}

func GetBooks(db *DB) ([]model.Book, error) {
	rows, err := db.Query(`SELECT * FROM books`)
	if err != nil {
		log.Printf("GetBooks: error getting books: %v", err)
		return nil, err
	}
	var books []model.Book
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublisedDate, &book.Description, &book.Cover, &book.CreatedAt, &book.UpdatedAt); err != nil {
			log.Printf("GetBooks: error scanning book: %v", err)
			return nil, err
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		log.Printf("GetBooks: rows error: %v", err)
		return nil, err
	}
	return books, err
}

func GetBookById(db *DB, id int) (*model.Book, error) {
	var book model.Book
	err := db.QueryRow(`SELECT * FROM books WHERE id = $1`, id).Scan(&book.ID, &book.Title, &book.Author, &book.PublisedDate, &book.Description, &book.Cover, &book.CreatedAt, &book.UpdatedAt)
	return &book, err
}

func UpdateBook(db *DB, id int, book *model.Book) error {
	currentBook, err := GetBookById(db, id)
	if err != nil {
		log.Printf("UpdateBook: error getting book by ID: %v", err)
		return err
	}

	if book.Title == "" {
		book.Title = currentBook.Title
	}
	if book.Author == "" {
		book.Author = currentBook.Author
	}
	if book.PublisedDate == "" {
		book.PublisedDate = currentBook.PublisedDate
	}
	if book.Description == "" {
		book.Description = currentBook.Description
	}
	if book.Cover == "" {
		book.Cover = currentBook.Cover
	}

	book.UpdatedAt = time.Now().Format(time.RFC3339)

	_, err = db.Exec(`UPDATE books SET title = $1, author = $2, published_date = $3, description = $4, cover = $5, updated_at =$6 WHERE id = $7`, book.Title, book.Author, book.PublisedDate, book.Description, book.Cover, book.UpdatedAt, id)
	return err
}
