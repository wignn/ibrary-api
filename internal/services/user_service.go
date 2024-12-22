package services

import (
	"github.com/wignn/library-api/internal/model"
	"github.com/wignn/library-api/internal/repository"
)

func GetUserById(db *repository.DB, id int) (*model.User, error){
	return repository.GetUserById(db, id)
}