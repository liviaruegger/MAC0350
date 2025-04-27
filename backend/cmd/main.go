package main

import (
	"strconv"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/liviaruegger/MAC0350/backend/internal/domain"
)

var profiles = []domain.User{
	{ID: 1, Name: "João da Silva", Email: "joao@example.com", City: "São Paulo", Phone: "(00) 0 0000-0000", Activities: activities_1},
	{ID: 2, Name: "Maria Souza", Email: "maria@example.com", City: "Não-Me-Toque", Phone: "(00) 0 0000-0000", Activities: activities_2},
	{ID: 3, Name: "Ana Costa", Email: "ana@example.com", City: "Vitória", Phone: "(00) 0 0000-0000", Activities: activities_3},
}

var activities_1 = []domain.Activity{}

var activities_2 = []domain.Activity{}

var activities_3 = []domain.Activity{}

func main() {
	router := gin.Default()
	router.GET("/profiles", getProfiles)
	router.GET("/profiles/:id", getProfileByID)
	router.POST("/profiles", postProfiles)

	router.Run("localhost:8080")
}

func getProfiles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, profiles)
}

func postProfiles(c *gin.Context) {
	var newProfile domain.User

	if err := c.BindJSON(&newProfile); err != nil {
		return
	}

	profiles = append(profiles, newProfile)
	c.IndentedJSON(http.StatusCreated, newProfile)
}

func getProfileByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range profiles {
		parsed, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			fmt.Println("Conversion error:", err)
			return
		}
		if a.ID == uint(parsed) {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "profile not found"})
}

