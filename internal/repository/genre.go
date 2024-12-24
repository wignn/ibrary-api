package repository

import (
	"log"

	"github.com/wignn/library-api/internal/model"
)

func CreateGenre(db *DB, genre *model.Genre) error {
	_, err := db.Exec("INSERT INTO genres (name) VALUES ($1)", genre.Name)
	return err
}

func GetGenres(db *DB) ([]model.Genre, error) {
	rows, err := db.Query("SELECT * FROM genres")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []model.Genre
	for rows.Next() {
		var genre model.Genre
		err := rows.Scan(&genre.ID, &genre.Name ,&genre.CreatedAt)
		if err != nil {
			log.Printf("GetGenres: error scanning genre: %v", err)
			return nil, err
		}
		genres = append(genres, genre)
	}
	return genres, nil
}
