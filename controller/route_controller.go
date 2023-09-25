package controller

import (
	"context"
	"log"
	"losangeles/config"
	"losangeles/service"
	"losangeles/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/"+strings.ToLower(config.Service.Name)+"/ping", Ping)
	router.GET("/tournaments", GetAllTournaments)
	router.GET("/tournaments/:id", GetTournamentByID)
	router.POST("/tournaments", CreateTournament)
	router.GET("/tournaments/:id/teams", GetAllTeamsForTournament)
	router.GET("/tournaments/:id/teams/:teamID", GetTeamForTournament)
	router.POST("/tournaments/:id/teams/:teamID", SetTeamForTournament)
	router.DELETE("/tournaments/:id/teams/:teamID", RemoveTeamFromTournament)
	router.GET("/teams/tournaments/:teamID", GetTournamentsForTeam)
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.SugarLogger.Infoln("GATEWAY REQUEST ID: " + c.GetHeader("Request-ID"))
		c.Next()
	}
}

func AuthChecker() gin.HandlerFunc {
	return func(c *gin.Context) {

		var requestUserID string

		ctx := context.Background()
		client, err := service.FirebaseAdmin.Auth(ctx)
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
		}
		if c.GetHeader("Authorization") != "" {
			token, err := client.VerifyIDToken(ctx, strings.Split(c.GetHeader("Authorization"), "Bearer ")[1])
			if err != nil {
				utils.SugarLogger.Errorln("error verifying ID token")
				requestUserID = "null"
			} else {
				utils.SugarLogger.Infoln("Decoded User ID: " + token.UID)
				requestUserID = token.UID
			}
		} else {
			utils.SugarLogger.Infoln("No user token provided")
			requestUserID = "null"
		}
		c.Set("userID", requestUserID)
		// The main authentication gateway per request path
		// The requesting user's ID and roles are pulled and used below
		// Any path can also be quickly halted if not ready for prod
		c.Next()
	}
}

func contains(s []string, element string) bool {
	for _, i := range s {
		if i == element {
			return true
		}
	}
	return false
}
