package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wignn/library-api/internal/repository"
	"github.com/joho/godotenv"
	"os"
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
	r.SetTrustedProxies([]string{"localhost"})
	r.Run(port)
}