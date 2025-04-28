package main

import (
	"github.com/gin-gonic/gin"

	"github.com/liviaruegger/MAC0350/backend/internal/handler"
)

func main() {
	router := gin.Default()
	router.GET("/profiles", handler.GetProfiles)
	router.GET("/profiles/:id", handler.GetProfileByID)
	router.POST("/profiles", handler.PostProfiles)

	router.Run("localhost:8080")
}