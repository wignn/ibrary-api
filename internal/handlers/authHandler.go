package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/wignn/library-api/internal/model"
	"github.com/wignn/library-api/internal/repository"
	"github.com/wignn/library-api/internal/services"
	"log"
	"net/http"
)

func LoginHandler(db *repository.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			log.Printf("LoginHandler: error binding JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid request"})
			return
		}

		token, err := services.LoginUser(db, user.Username, user.Password)
		if err != nil {
			log.Printf("LoginHandler: error logging in user: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"errors": "Credentials error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": token})
	}
}

func RegisterHandler(db *repository.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			log.Printf("RegisterHandler: error binding JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"errors": "Internal server error"})
			return
		}
		err := services.RegisterRequest(db, &user)
		if err != nil {
			log.Printf("RegisterHandler: error registering user: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"errors": "user already exists"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	}
}
