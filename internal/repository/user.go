package repository

import (
	"database/sql"
	"log"
	"time"

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
	return &user, err
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
	user.UpdatedAt = time.Now().Format(time.RFC3339)
	_, err = db.Exec(`UPDATE users SET username = $1, profile_picture = $2, email = $3, updated_at=$4 WHERE id = $5`,
		user.Username, user.ProfilePicture, user.Email, user.UpdatedAt, user.ID)
	return user, err
}

func UpdateUserToken(db *DB, user *model.User) error {
	log.Println("Updating user token ", *user.Token, user.ID)
	_, err := db.Exec(`UPDATE users SET token = $1 WHERE id = $2`, *user.Token, user.ID)
	return err
}

func ResetPassword(db *DB, id int, hashedPassword, token string) error {
	var user model.User
	err := db.QueryRow(`SELECT id FROM users WHERE id = $1 AND token = $2`, id, token).Scan(&user.ID, &user.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No user found")
			return err
		} else {
			log.Panicf("Query error: %v", err)
			return err
		}
	}
	_, err = db.Exec(`UPDATE users SET password = $1, token= $2 WHERE id = $3`, hashedPassword, nil, id)
	return err
}
