package repository

import (
	"database/sql"

	"github.com/wignn/library-api/internal/model"
)

func CreateUser(db *DB, user *model.User) error {
	_, err := db.Exec(`INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`, user.Username, user.Email, user.Password)
	return err
}

func GetUserById(db *DB, id int) (*model.User, error) {
	var user model.User
	var profilePicture sql.NullString
	err := db.QueryRow(`SELECT * FROM users WHERE id = $1`, id).Scan(&user.ID, &user.Username, &user.Email, &profilePicture, &user.CreatedAt, &user.UpdatedAt, )
	if err != nil {
		return nil, err
	}
	if profilePicture.Valid {
		user.ProfilePicture = &profilePicture.String
	}else {
		user.ProfilePicture = nil
	}

	return &user, nil
}