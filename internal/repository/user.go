package repository

import (
	"database/sql"
	"fmt"

	"github.com/wignn/library-api/internal/model"
)

func CreateUser(db *DB, user *model.User) error {
	_, err := db.Exec(`INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`, user.Username, user.Email, user.Password)
	return err
}

func GetUserById(db *DB, id int) (*model.GetUserResponse, error) {
	var user model.GetUserResponse
	var profilePicture sql.NullString
	err := db.QueryRow(`SELECT id, username, email, profile_picture, created_at, updated_at FROM users WHERE id = $1`, id).Scan(&user.ID, &user.Username, &user.Email, &profilePicture, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if profilePicture.Valid {
		user.ProfilePicture = &profilePicture.String
	} else {
		user.ProfilePicture = nil
	}
	return &user, nil
}

func GetUserByUsername(db *DB, username string) (*model.User, error) {
	var user model.User
	err := db.QueryRow(`SELECT id, username, email, profile_picture, password, is_active, is_admin, created_at, updated_at FROM users WHERE username = $1`, username).Scan(
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

func UpdateUserProfile(db *DB, user *model.User) (*model.User, error) {
	var currentUser *model.GetUserResponse
	currentUser, err := GetUserById(db, user.ID)
	if err != nil {
		return nil, err
	}

	if user.Username == "" {
		user.Username = currentUser.Username
	}
	if user.ProfilePicture == nil || *user.ProfilePicture == "" {
		user.ProfilePicture = currentUser.ProfilePicture
	}
	if user.Email == "" {
		user.Email = currentUser.Email
	}

	_, err = db.Exec(`UPDATE users SET username = $1, profile_picture = $2, email = $3 WHERE id = $4`,
		user.Username, user.ProfilePicture, user.Email, user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUserToken(db *DB, id int, token string) error {
	_, err := db.Exec(`UPDATE users SET token = $1 WHERE id = $2`, token, id)
	return err
}

func ResetPassword(db *DB, id int, hashedPassword, token string) error {
	var user model.User
	err := db.QueryRow(`SELECT id FROM users WHERE id = $1 AND token = $2`, id, token).Scan(&user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("User not found")
		}
		return err
	}

	_, err = db.Exec(`UPDATE users SET password = $1 WHERE id = $2`, hashedPassword, id)
	return err
}
