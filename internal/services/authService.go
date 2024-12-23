package services

import (
	"github.com/wignn/library-api/internal/model"
	"github.com/wignn/library-api/internal/repository"
	"github.com/wignn/library-api/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func RegisterRequest(db *repository.DB, user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return repository.CreateUser(db, user)
}

func LoginUser(db *repository.DB, username, password string) (string, error) {
	user, err := repository.GetUserByUsername(db, username)
	if err != nil {
		return "", err
	}
	
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}
	
	return utils.GenerateToken(user.Username, user.ID, user.IsAdmin)
}
