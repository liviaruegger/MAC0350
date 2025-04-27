package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type profile struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	City       string     `json:"city"`
	Phone      string     `json:"phone"`
	Activities []activity `json:"activities"`
}

type activity struct {
	ID       string    `json:"id"`
	Distance string    `json:"distance"`
	Start    time.Time `json:"start"`
	Finish   time.Time `json:"finish"`
	Size     time.Time `json:"size"`
	Laps     int       `json:"laps"`
}

var profiles = []profile{
	{ID: "1", Name: "João da Silva", Email: "joao@example.com", City: "São Paulo", Phone: "(00) 0 0000-0000", Activities: activities_1},
	{ID: "2", Name: "Maria Souza", Email: "maria@example.com", City: "Não-Me-Toque", Phone: "(00) 0 0000-0000", Activities: activities_2},
	{ID: "3", Name: "Ana Costa", Email: "ana@example.com", City: "Vitória", Phone: "(00) 0 0000-0000", Activities: activities_3},
}

var activities_1 = []activity{}

var activities_2 = []activity{}

var activities_3 = []activity{}

func main() {
	router := gin.Default()
	router.GET("/profiles", getProfiles)
	router.GET("/profiles/:id", getProfileById)
	router.POST("/profiles", postProfiles)

	router.Run("localhost:8080")
}

func getProfiles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, profiles)
}

func getProfileByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range profiles {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "profile not found"})
}

func postProfiles(c *gin.Context) {
	var newProfile profile

	if err := c.BindJSON(&newProfile); err != nil {
		return
	}

	profiles = append(profiles, newProfile)
	c.IndentedJSON(http.StatusCreated, newProfile)
}
