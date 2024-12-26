package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/wignn/library-api/internal/model"
)

func CreateUser(db *DB, user *model.User) error {
	stmt, err := db.Prepare(`INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`)
	if err != nil {
		log.Println("error preparing statement: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, user.Password)
	if err != nil {
		log.Println("error creating user: ", err)
		return fmt.Errorf("error creating user")
	}
	return err
}

func GetUserById(db *DB, id int) (*model.GetUserResponse, error) {
	stmt, err := db.Prepare(`SELECT id, username, email, profile_picture, created_at, updated_at FROM users WHERE id = $1`)
	if err != nil {
		log.Println("error preparing statement: ", err)
		return nil, err
	}
	defer stmt.Close()

	var user model.GetUserResponse
	var profilePicture sql.NullString
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Username, &user.Email, &profilePicture, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	if profilePicture.Valid {
		user.ProfilePicture = &profilePicture.String
	} else {
		user.ProfilePicture = nil
	}

	return &user, err
}

func GetUserByUsername(db *DB, username string) (*model.User, error) {
	stmt, err := db.Prepare(`SELECT id, username, email, profile_picture, password, is_active, is_admin, created_at, updated_at FROM users WHERE username = $1`)
	if err != nil {
		log.Println("error preparing statement: ", err)
		return nil, err
	}
	defer stmt.Close()

	var user model.User
	err = stmt.QueryRow(username).Scan(
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
	stmt, err := db.Prepare(`UPDATE users SET username = $1, profile_picture = $2, email = $3, updated_at = $4 WHERE id = $5`)
	if err != nil {
		log.Println("error preparing statement: ", err)
		return nil, err
	}
	defer stmt.Close()

	var currentUser *model.GetUserResponse
	currentUser, err = GetUserById(db, user.ID)
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

	_, err = stmt.Exec(user.Username, user.ProfilePicture, user.Email, user.UpdatedAt, user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUserToken(db *DB, user *model.User) error {
	stmt, err := db.Prepare(`UPDATE users SET token = $1 WHERE id = $2`)
	if err != nil {
		log.Println("error preparing statement: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(*user.Token, user.ID)
	if err != nil {
		log.Println("error updating user token: ", err)
	}
	return err
}

func ResetPassword(db *DB, id int, hashedPassword, token string) error {
	stmtSelect, err := db.Prepare(`SELECT id FROM users WHERE id = $1 AND token = $2`)
	if err != nil {
		log.Println("error preparing select statement: ", err)
		return err
	}
	defer stmtSelect.Close()

	var user model.User
	err = stmtSelect.QueryRow(id, token).Scan(&user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No user found")
			return err
		} else {
			log.Panicf("Query error: %v", err)
			return err
		}
	}

	stmtUpdate, err := db.Prepare(`UPDATE users SET password = $1, token = $2 WHERE id = $3`)
	if err != nil {
		log.Println("error preparing update statement: ", err)
		return err
	}
	defer stmtUpdate.Close()

	_, err = stmtUpdate.Exec(hashedPassword, nil, id)
	if err != nil {
		log.Println("error resetting password: ", err)
	}
	return err
}
