package repository

import (
	"github.com/wignn/library-api/internal/model"
	"log"
	"time"
)

func CreateBook(db *DB, book *model.Book) error {
	stmt, err := db.Prepare(`INSERT INTO books (title, author, published_date, description, cover) VALUES ($1, $2, $3, $4, $5)`)
	if err != nil {
		log.Printf("CreateBook: error preparing statement: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(book.Title, book.Author, book.PublisedDate, book.Description, book.Cover)
	return err
}

func GetBooks(db *DB) ([]model.Book, error) {
	stmt, err := db.Prepare(`SELECT * FROM books`)
	if err != nil {
		log.Printf("GetBooks: error preparing statement: %v", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Printf("GetBooks: error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublisedDate, &book.Description, &book.Cover, &book.CreatedAt, &book.UpdatedAt); err != nil {
			log.Printf("GetBooks: error scanning book: %v", err)
			return nil, err
		}
		books = append(books, book)
	}
	return books, err
}

func GetBookById(db *DB, id int) (*model.Book, error) {
	stmt, err := db.Prepare("SELECT * FROM books WHERE id = $1")
	if err != nil {
		log.Printf("GetBookById: error preparing statement: %v", err)
		return nil, err
	}
	defer stmt.Close()
	var book model.Book
	err = stmt.QueryRow(id).Scan(&book.ID, &book.Title, &book.Author, &book.PublisedDate, &book.Description, &book.Cover, &book.CreatedAt, &book.UpdatedAt)
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

	stmt, err := db.Prepare(`UPDATE books SET title = $1, author = $2, published_date = $3, description = $4, cover = $5, updated_at =$6 WHERE id = $7`)
	if err != nil {
		log.Printf("UpdateBook: error preparing statement: %v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(book.Title, book.Author, book.PublisedDate, book.Description, book.Cover, book.UpdatedAt, id)
	return err
}

func GetBookByName(db *DB, name string) (*model.Book, error) {
	stmt, err := db.Prepare(`SELECT * FROM books WHERE title = $1`)
	if err != nil {
		log.Printf("GetBookByNaem: error preparing statement: %v", err)
		return nil, err
	}
	defer stmt.Close()

	var book model.Book
	err = stmt.QueryRow(name).Scan(&book.ID, &book.Title, &book.Author, &book.PublisedDate, &book.Description, &book.Cover, &book.CreatedAt, &book.UpdatedAt)
	return &book, err
}

func DeleteBook(db *DB, id int) error {
	stmt, err := db.Prepare(`DELETE FROM books WHERE id = $1`)
	if err != nil {
		log.Printf("DeleteBook: error preparing statement: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
