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
	err := db.QueryRow(`SELECT id, username, email, profile_picture, created_at, updated_at FROM users WHERE id = $1`, id).Scan(&user.ID, &user.Username, &user.Email, &profilePicture, &user.CreatedAt, &user.UpdatedAt, )
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


func GetUserByUsername(db *DB, username string) (*model.User, error) {
	var user model.User
	err := db.QueryRow(`SELECT id, username, email, profile_picture, is_active, is_admin, created_at, updated_at FROM users WHERE username = $1`, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.ProfilePicture,
		&user.Password,
		&user.IsActive,
		&user.IsAdmin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
