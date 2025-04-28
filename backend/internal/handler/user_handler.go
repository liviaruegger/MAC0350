package handler

import (
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/liviaruegger/MAC0350/backend/internal/repository"

	"net/http"
	"strconv"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetProfiles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, repository.Profiles)
}

func PostProfiles(c *gin.Context) {
	var newProfile domain.User

	if err := c.BindJSON(&newProfile); err != nil {
		return
	}

	repository.Profiles = append(repository.Profiles, newProfile)
	c.IndentedJSON(http.StatusCreated, newProfile)
}

func GetProfileByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range repository.Profiles {
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