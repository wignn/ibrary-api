package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wignn/library-api/internal/model"
	"github.com/wignn/library-api/internal/repository"
	"github.com/wignn/library-api/internal/services"
)

func CreateGenreHandler(db *repository.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book model.Genre
		if err := c.ShouldBindBodyWithJSON(&book); err != nil {
			c.JSON(400, gin.H{"errors": "Invalid request"})
			return
		}
		err := services.CreateGenre(db, &book)
		if err != nil {
			log.Printf("eror create genre: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Genre created successfully"})
	}
}

func GetGenreListHandler(db *repository.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		genres, err := services.GetGenres(db)
		if err != nil {
			log.Printf("error listing genres: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(http.StatusOK, genres)
	}
}

func GetGenreByIdHandler(db *repository.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			log.Printf("error converting ID: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		genre, err := services.GetGenreById(db, id)
		if err != nil {
			log.Printf("error getting genre: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, genre)
	}
}

func UpdateGenreHandler(db *repository.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			log.Printf("error converting ID: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var genre model.Genre
		if err := c.ShouldBindBodyWithJSON(&genre); err != nil {
			log.Printf("error binding JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		err = services.UpdateGenre(db, id, &genre)
		if err != nil {
			log.Printf("error updating genre: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Genre updated successfully"})
	}

}

func DeleteGenreHandler(db *repository.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			log.Printf("error converting ID: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "validation error"})
			return
		}

		err = services.DeleteGenre(db, id)

		if err != nil {
			log.Printf("error deleting genre: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Genre deleted successfully"})
	}
}

func AddGenreToBookHandler(db *repository.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var genre model.BookGenre
		if err := c.ShouldBindBodyWithJSON(&genre); err != nil {
			log.Printf("AddGenreToBookHandler: error binding JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid request"})
			return
		}
		err := services.AddGenreToBook(db, &genre)
		if err != nil {
			log.Printf("AddGenreToBookHandler: error adding genre to book: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"errors": "Internal server error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Genre added to book successfully"})
	}
}

// func GetBooksByGenreBookHandler(db *repository.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		id, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			log.Printf("GetBooksByGenreHandler: error converting ID: %v", err)
// 			c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid ID"})
// 			return
// 		}
// 		books, err := services.GetBooksByGenreBook(db, id)
// 		if err != nil {
// 			log.Printf("GetBooksByGenreHandler: error getting books by genre: %v", err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"errors": "Internal server error"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, books)
// 	}
// }