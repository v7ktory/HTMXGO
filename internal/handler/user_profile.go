package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/v7ktory/htmx+go/internal/model"
)

func (h *Handler) UserProfile(c *gin.Context) {

	userSession := c.MustGet("userSession").(*model.UserSession)

	// Отображение страницы профиля пользователя
	c.HTML(http.StatusOK, "profile.html", gin.H{
		"ID":    userSession.ID,
		"Name":  userSession.Name,
		"Email": userSession.Email,
	})
}
