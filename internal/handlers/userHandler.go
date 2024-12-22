package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/wignn/library-api/internal/repository"
	"github.com/wignn/library-api/internal/services"
)

func GetUserByIdHandler(db *repository.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}

		user, err := services.GetUserById(db, id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUserHandler(db *repository.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}

		var userUpdate struct {
			Username string `json:"username"`
			Email    string `json:"email"`
		}
		fmt.Println(id)

		if err := c.ShouldBindJSON(&userUpdate); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request payload"})
			return
		}

	}
}