package handler

import "github.com/gin-gonic/gin"

func (h *Handler) SignOut(c *gin.Context) {

	// Get the session from the authorization header
	sessionHeader := c.GetHeader("Authorization")

	// Ensure the session header is not empty and in the correct format
	if sessionHeader == "" || len(sessionHeader) < 8 || sessionHeader[:7] != "Bearer " {
		handleError(c, "invalid session header", 400)
		return
	}

	// Get the session id
	sessionID := sessionHeader[7:]

	// Delete the session
	err := h.Service.SignOut(sessionID)
	if err != nil {
		handleError(c, "failed to sign out", 500)
		return
	}

	c.JSON(200, gin.H{"success": true})
}
