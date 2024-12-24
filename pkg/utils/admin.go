package utils

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
)

func IsAdmin(c *gin.Context)  error {
	isAdminValue, exists := c.Get("isAdmin")
	if !exists {
		log.Printf("isAdmin value not found")
		return fmt.Errorf("isAdmin value not found")
	}

	isAdmin, ok := isAdminValue.(bool)
	if !ok || !isAdmin {
		log.Printf("User is not an admin")
		return fmt.Errorf("user is not an admin")
	}

	return nil
}