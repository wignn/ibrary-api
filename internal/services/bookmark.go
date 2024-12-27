package services

import (
	"github.com/wignn/library-api/internal/model"
	"github.com/wignn/library-api/internal/repository"
)

func CreateBookmark(db *repository.DB, bookmark *model.Bookmark) error {
	return repository.CreateBookmark(db, bookmark)
}

func DeleteBookmark(db *repository.DB, id string) error {
	return repository.DeleteBookmark(db, id)
}