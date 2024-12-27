package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wignn/library-api/internal/model"
	"github.com/wignn/library-api/internal/repository"
	"github.com/wignn/library-api/internal/services"
)

func CreateBookmarkHandler(db *repository.DB)gin.HandlerFunc{
	return func(c *gin.Context){
		var bookmarks model.Bookmark

		if err := c.ShouldBindBodyWithJSON(&bookmarks); err != nil {
			log.Printf("CreateBookmarkHandler: error binding JSON: %v", err)
			c.JSON(400, gin.H{"errors": "Invalid request"})
			return
		}

		err := services.CreateBookmark(db ,&bookmarks)

		if err != nil {
			log.Printf("CreateBookmarkHandler: error creating bookmark: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"errors": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Bookmark created successfully"})
	}	
}

func DeleteBookmarkHandler(db *repository.DB)gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")
		err := services.DeleteBookmark(db, id)
		if err != nil {
			log.Printf("DeleteBookmarkHandler: error deleting bookmark: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"errors": "Internal server error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Bookmark deleted successfully"})
	}
}

func GetBookmarkListHandler(db *repository.DB){
	
}