package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wignn/library-api/internal/auth"
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
		apiV1.PATCH("/users/:id/verify-email", handlers.VerifyEmailHandler(db))
		apiV1.PUT("/users/:id/reset-password", handlers.ResetPasswordHandler(db))
		aunthenticated := apiV1.Group("/")
		{
			//user route
			aunthenticated.Use(auth.AuthMIddleware())
			aunthenticated.GET("/users/:id", handlers.GetUserByIdHandler(db))
			aunthenticated.PUT("/users/:id", handlers.UpdateUserHandler(db))

			//book route
			aunthenticated.POST("/books", handlers.CreateBooksHandler(db))
			aunthenticated.GET("/books", handlers.GetBookListHandler(db))
			aunthenticated.GET("/books/:id", handlers.GetBookByIdHandler(db))
			aunthenticated.PUT("/books/:id", handlers.UpdateBookHandler(db))
			aunthenticated.DELETE("/books/:id", handlers.DeleteBookHandler(db))

			//genre route
			aunthenticated.POST("/genres", handlers.CreateGenreHandler(db))
			aunthenticated.GET("/genres", handlers.GetGenreListHandler(db))
			aunthenticated.GET("/genres/:id", handlers.GetGenreByIdHandler(db))
			aunthenticated.PUT("/genres/:id", handlers.UpdateGenreHandler(db))
			aunthenticated.DELETE("/genres/:id", handlers.DeleteGenreHandler(db))	

		}
	}
}
