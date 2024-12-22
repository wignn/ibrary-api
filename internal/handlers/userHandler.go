package handlers

import (
	"log"
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
			log.Printf("Error converting ID: %v", err)
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}

		user, err := services.GetUserById(db, id)
		if err != nil {
			log.Printf("Error getting user by ID: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUserHandler(db *repository.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("Error converting user ID: %v", err)
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}

		var userUpdate struct {
			Username       string `json:"username"`
			Email          string `json:"email"`
			ProfilePicture string `json:"profile_picture"`
		}

		if err := c.ShouldBindJSON(&userUpdate); err != nil {
			log.Printf("Error binding JSON: %v", err)
			c.JSON(400, gin.H{"error": "Invalid request payload"})
			return
		}

		_, err = services.UpdateUserProfile(db, id, userUpdate.Username, userUpdate.ProfilePicture, userUpdate.Email)
		if err != nil {
			log.Printf("Error updating user profile: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User profile updated successfully"})
	}
}


func VerifyEmailHandler(db *repository.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("Error converting ID: %v", err)
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}

		err = services.SendEmailVerification(db, id)
		if err != nil {
			log.Printf("Error sending email verification: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
	}
}