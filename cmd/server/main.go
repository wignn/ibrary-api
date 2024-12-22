package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wignn/library-api/internal/repository"
	"github.com/wignn/library-api/internal/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	db, err := repository.InitDb()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	r := gin.Default()

	routes.InitRoutes(r, db)

	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}