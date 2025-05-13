// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liviaruegger/MAC0350/backend/config"
	"github.com/liviaruegger/MAC0350/backend/internal/app"
	"github.com/liviaruegger/MAC0350/backend/internal/handler"
	"github.com/liviaruegger/MAC0350/backend/internal/repository"
)

func SetupRouter() *gin.Engine {
	db := config.SetupDatabase()

	userRepo := repository.NewUserRepository(db)
	userService := app.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	router.GET("/users/:id", userHandler.GetUserByID)
	router.POST("/users", userHandler.CreateUser)

	return router
}

func main() {
	router := SetupRouter()
	router.Run(":8080")
}
