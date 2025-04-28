package main

import (
	"github.com/gin-gonic/gin"

	"github.com/liviaruegger/MAC0350/backend/internal/config"
	"github.com/liviaruegger/MAC0350/backend/internal/repository"
	"github.com/liviaruegger/MAC0350/backend/internal/app"
	"github.com/liviaruegger/MAC0350/backend/internal/handler"
)

func main() {
	db := config.SetupDatabase()

    userRepo := repository.NewUserRepository(db)
    userService := app.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	router.GET("/users/:id", userHandler.GetUserByID)
	router.POST("/users", userHandler.CreateUser)

	router.Run("localhost:8080")
}