package repository

import (
	"fmt"
	"log"

	"github.com/wignn/library-api/internal/model"
)

func CreateGenre(db *DB, genre *model.Genre) error {
	existingGenre, err := GetGenreByName(db, genre.Name)
	if err == nil && existingGenre != nil {
		return fmt.Errorf("genre already exists")
	}
	_, err = db.Exec("INSERT INTO genres (name) VALUES ($1)", genre.Name)
	log.Printf("error creating genre: %v",err )
	return fmt.Errorf("error creating genre")
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
		err := rows.Scan(&genre.ID, &genre.Name, &genre.CreatedAt)
		if err != nil {
			log.Printf("GetGenres: error scanning genre: %v", err)
			return nil, err
		}
		genres = append(genres, genre)
	}
	return genres, nil
}

func GetGenreById(db *DB, id int) (*model.Genre, error) {
	var genre model.Genre
	err := db.QueryRow("SELECT * FROM genres WHERE id = $1", id).Scan(&genre.ID, &genre.Name, &genre.CreatedAt)
	return &genre, err
}

func UpdateGenre(db *DB, id int, genre *model.Genre) error {
	_, err := db.Exec("UPDATE genres SET name = $1 WHERE id = $2", genre.Name, id)
	return err
}


func GetGenreByName(db *DB, name string) (*model.Genre, error) {
	var genre model.Genre
	err := db.QueryRow("SELECT * FROM genres WHERE name = $1", name).Scan(&genre.ID, &genre.Name, &genre.CreatedAt)
	return &genre, err
}