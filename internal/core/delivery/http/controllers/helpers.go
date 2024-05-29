package controllers

import (
	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{"error": err.Error()})
}
