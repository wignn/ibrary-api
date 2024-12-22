package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/wignn/library-api/internal/model"
	"github.com/wignn/library-api/internal/repository"
	"github.com/wignn/library-api/internal/services"
)

func LoginHandler(db *repository.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User;
		fmt.Println(user)
	}
}

func RegisterHandler(db *repository.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := services.RegisterRequest(db, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	}
}

