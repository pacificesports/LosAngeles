package controller

import (
	"github.com/gin-gonic/gin"
	"losangeles/model"
	"losangeles/service"
	"net/http"
)

func GetAllTeamsForTournament(c *gin.Context) {
	result := service.GetAllTeamsForTournament(c.Param("id"))
	c.JSON(http.StatusOK, result)
}

func GetTeamForTournament(c *gin.Context) {
	result := service.GetTeamForTournament(c.Param("id"), c.Param("teamID"))
	if result.TeamID == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "Failed to find team with id: " + c.Param("teamID") + " in tournament with id: " + c.Param("tournamentID")})
		return
	}
	c.JSON(http.StatusOK, result)
}

func SetTeamForTournament(c *gin.Context) {
	var input model.TournamentTeam
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if service.GetTeamForTournament(c.Param("id"), input.TeamID).TeamID != "" {
		c.JSON(http.StatusConflict, gin.H{"message": "Team already in tournament"})
		return
	}
	if err := service.SetTeamForTournament(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service.GetTeamForTournament(c.Param("id"), input.TeamID))
}

func RemoveTeamFromTournament(c *gin.Context) {
	if err := service.RemoveTeamFromTournament(c.Param("id"), c.Param("teamID")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Team removed from tournament"})
}

func GetTournamentsForTeam(c *gin.Context) {
	result := service.GetTournamentsForTeam(c.Param("teamID"))
	c.JSON(http.StatusOK, result)
}
