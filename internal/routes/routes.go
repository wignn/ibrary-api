package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wignn/library-api/internal/handlers"
	"github.com/wignn/library-api/internal/repository"
)

func InitRoutes(r *gin.Engine, db *repository.DB) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "api working",
		})
	})

	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/login", handlers.LoginHandler(db))
		apiV1.POST("/register", handlers.RegisterHandler(db))
		// apiV1.GET("/users/:id", handlers.GetUserByIdHandler(db))
	}
}