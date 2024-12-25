package handlers

import (
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/wignn/library-api/internal/model"
	"github.com/wignn/library-api/internal/repository"
	"github.com/wignn/library-api/internal/services"
	"github.com/wignn/library-api/pkg/utils"
)

func CreateBooksHandler(db *repository.DB) gin.HandlerFunc{
	return func (c *gin.Context){
		err := utils.IsAdmin(c)
		if err != nil {
			log.Printf("GetBookByIdHandler: error checking admin: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		var book model.Book
		if err  := c.ShouldBindBodyWithJSON(&book); err != nil {
			log.Printf("GetBooksHandler: error binding JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid request"})
			return
		}

		err = services.CreateBook(db, &book)
		if err != nil {
			log.Printf("GetBooksHandler: error creating book: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"errors": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book created successfully"})
	}
}

func GetBookListHandler(db *repository.DB) gin.HandlerFunc{
	return func (c *gin.Context){
		books, err := services.GetBooks(db)
		if err != nil {
			log.Printf("GetBooksHandler: error getting books: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"errors": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, books)
	}
}


func GetBookByIdHandler(db *repository.DB) gin.HandlerFunc{
	return func (c *gin.Context){
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("GetBookByIdHandler: error converting ID: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid ID"})
			return
		}

		book, err := services.GetBookById(db, id)
		if err != nil {
			log.Printf("GetBookByIdHandler: error getting book by ID: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"errors": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, book)
	}
}

func UpdateBookHandler(db *repository.DB) gin.HandlerFunc{
	return func (c *gin.Context){
		err := utils.IsAdmin(c)
		if err != nil {
			log.Printf("UpdateBookHandler: error checking admin: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Unauthorized"})
			return
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			log.Printf("UpdateBookHandler: error converting ID: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid ID"})
			return
		}

		var book model.Book
		
		if err := c.ShouldBindBodyWithJSON(&book); err != nil {
			log.Printf("UpdateBookHandler: error binding JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid request"})
			return
		}

		err = services.UpdateBook(db, id, &book)
		
		if err != nil {
			log.Printf("UpdateBookHandler: error updating book: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"errors": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
	}
}

func DeleteBookHandler(db *repository.DB) gin.HandlerFunc{
	return func (c *gin.Context){
		err := utils.IsAdmin(c)
		if err != nil {
			log.Printf("DeleteBookHandler: error checking admin: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Unauthorized"})
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("DeleteBookHandler: error converting ID: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid ID"})
			return
		}

		err = services.DeleteBook(db, id)
		if err != nil {
			log.Printf("DeleteBookHandler: error deleting book: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"errors": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
	}
}