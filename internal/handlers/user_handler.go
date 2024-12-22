package handlers

import (
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
			c.JSON(http., gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, user)
	}
}