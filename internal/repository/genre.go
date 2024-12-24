package repository

import (
	"github.com/wignn/library-api/internal/model"
)

func CreateGenre(db *DB, genre *model.Genre) error {
	_, err := db.Exec("INSERT INTO genres (name) VALUES ($1)", genre.Name)
	return err
}