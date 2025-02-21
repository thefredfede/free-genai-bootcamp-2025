package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResetHistory(c *gin.Context) {
	// Implement the logic to reset history
	c.JSON(http.StatusOK, gin.H{"message": "ResetHistory"})
}

func FullReset(c *gin.Context) {
	// Implement the logic to perform a full reset
	c.JSON(http.StatusOK, gin.H{"message": "FullReset"})
}
