package controller

import (
	"losangeles/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": config.Service.Name + " v" + config.Version + " is online!"})
}
