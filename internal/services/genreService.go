package services

import (
	"github.com/wignn/library-api/internal/model"
	"github.com/wignn/library-api/internal/repository"
)

func CreateGenre(db *repository.DB, genre *model.Genre) error {
	return repository.CreateGenre(db, genre)

}

func GetGenres(db *repository.DB) ([]model.Genre, error) {
	return repository.GetGenres(db)
}

func GetGenreById(db *repository.DB, id int) (*model.Genre, error) {
	return repository.GetGenreById(db, id)
}

func UpdateGenre(db *repository.DB, id int, genre *model.Genre) error {
	return repository.UpdateGenre(db, id, genre)
}

func DeleteGenre(db *repository.DB, id int) error {
	return repository.DeleteGenre(db, id)
}
