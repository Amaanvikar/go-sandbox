package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHealth is a simple check that the server is running.
func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
