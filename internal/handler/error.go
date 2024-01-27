package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func handleError(c *gin.Context, errMsg string, status int) {
	c.JSON(status, gin.H{"error": errors.New(errMsg).Error()})
}
