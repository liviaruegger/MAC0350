package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liviaruegger/MAC0350/backend/config"
	"github.com/liviaruegger/MAC0350/backend/internal/app"
	"github.com/liviaruegger/MAC0350/backend/internal/handler"
	"github.com/liviaruegger/MAC0350/backend/internal/repository"

	// Swagger imports
	_ "github.com/liviaruegger/MAC0350/docs"
	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

// @title           Swim Tracker API
// @version         1.0
// @description     This is an API for tracking swimming
// @host            localhost:8080
// @BasePath        /

func SetupRouter() *gin.Engine {
	db := config.SetupDatabase()

	userRepo := repository.NewUserRepository(db)
	userService := app.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	intervalRepo := repository.NewIntervalRepository(db)
	intervalService := app.NewIntervalService(intervalRepo)
	intervalHandler := handler.NewIntervalHandler(intervalService)

	router := gin.Default()

	// Swagger route
	router.GET("/swagger/*any", ginswagger.WrapHandler(swaggerfiles.Handler))

	// User routes
	router.GET("/users/:id", userHandler.GetUserByID)
	router.POST("/users", userHandler.CreateUser)

	// Interval routes
	router.POST("/intervals", intervalHandler.CreateInterval)

	return router
}

func main() {
	router := SetupRouter()
	router.Run(":8080")
}
