package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wignn/library-api/internal/model"
	"github.com/wignn/library-api/internal/repository"
	"github.com/wignn/library-api/internal/services"
)


func CreateGenreHandler(db *repository.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book model.Genre
		if err := c.ShouldBindBodyWithJSON(&book); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}
		err := services.CreateGenre(db, &book)
		if err != nil {
			log.Printf("eror create genre", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Genre created successfully"})
	}
}

func GetGenreListHandler(db *repository.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		genres, err := services.GetGenres(db)
		if err != nil {
			log.Printf("eror list genre", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(http.StatusOK, genres)
	}
}

func GetGenreByIdHandler(db *repository.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		
	}
}