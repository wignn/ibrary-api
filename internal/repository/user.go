package repository

import "github.com/wignn/library-api/internal/model"

func CreateUser(db *DB, user *model.User) error {
	_, err := db.Exec(`INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`, user.Username, user.Email, user.Password)
	return err
}