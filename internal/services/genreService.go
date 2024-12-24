package services

import (
	"github.com/wignn/library-api/internal/model"
	"github.com/wignn/library-api/internal/repository"
)

func CreateGenre(db *repository.DB, genre *model.Genre) error {
	return repository.CreateGenre(db, genre)

}