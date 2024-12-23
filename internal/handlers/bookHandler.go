package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wignn/library-api/internal/model"
	"github.com/wignn/library-api/internal/repository"
)

func GetBooksHandler(db *repository.DB) gin.HandlerFunc{
	return func (c *gin.Context){
		var book model.Book
		if err  := c.ShouldBindBodyWithJSON(&book); err != nil {
			log.Printf("GetBooksHandler: error binding JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		}
	}

}