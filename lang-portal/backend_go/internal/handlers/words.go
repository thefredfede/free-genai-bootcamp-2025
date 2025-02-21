package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWords(c *gin.Context) {
	// Implement the logic to get words
	c.JSON(http.StatusOK, gin.H{"message": "GetWords"})
}

func GetWordByID(c *gin.Context) {
	// Implement the logic to get a word by ID
	c.JSON(http.StatusOK, gin.H{"message": "GetWordByID"})
}
