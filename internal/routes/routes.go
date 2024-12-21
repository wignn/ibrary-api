package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wignn/library-api/internal/handlers"
	"github.com/wignn/library-api/internal/repository"
)

func InitRoutes(r *gin.Engine, db *repository.DB) {
	//auth route
	r.POST("/login", handlers.LoginHandler(db))	
	r.POST("/register", handlers.RegisterHandler(db))	
}