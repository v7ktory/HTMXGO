package handler

import (
	"github.com/gin-gonic/gin"
)

func handleError(c *gin.Context, errMsg string, status int) {
	c.String(status, errMsg)
}
