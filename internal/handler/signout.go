package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignOut(c *gin.Context) {
	cookie, err := c.Cookie("sessionID")
	if err != nil || cookie == "" {
		handleError(c, "invalid session ID in Cookie", http.StatusBadRequest)
		return
	}

	if err := h.Service.SignOut(c, cookie); err != nil {
		handleError(c, "failed to sign out", http.StatusInternalServerError)
		return
	}

	c.Header("HX-Redirect", "/")
	c.Status(http.StatusOK)
}
