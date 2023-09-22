package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"losangeles/model"
	"losangeles/service"
	"net/http"
)

func GetAllTournaments(c *gin.Context) {
	result := service.GetAllTournaments()
	c.JSON(http.StatusOK, result)
}

func GetTournamentByID(c *gin.Context) {
	result := service.GetTournamentByID(c.Param("id"))
	if result.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "Failed to find tournament with id: " + c.Param("id")})
		return
	}
	c.JSON(http.StatusOK, result)
}

func CreateTournament(c *gin.Context) {
	var input model.Tournament
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if input.ID == "" {
		input.ID = uuid.New().String()
	}
	if result := service.CreateTournament(input); result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": result.Error()})
		return
	}
	c.JSON(http.StatusOK, service.GetTournamentByID(input.ID))
}
