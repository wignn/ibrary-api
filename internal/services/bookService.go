package services

import (
	"github.com/wignn/library-api/internal/model"
	"github.com/wignn/library-api/internal/repository"
)

func CreateBook(db *repository.DB, book *model.Book) error {
	return repository.CreateBook(db, book)
}

func GetBooks(db *repository.DB) ([]model.Book, error) {
	return repository.GetBooks(db)
}

func GetBookById(db *repository.DB, id int) (*model.Book, error) {
	return repository.GetBookById(db, id)
}

func UpdateBook(db *repository.DB,id int, book *model.Book) error {
	return repository.UpdateBook(db, id, book)
}

func DeleteBook(db *repository.DB, id int) error {
	return repository.DeleteBook(db, id)
}