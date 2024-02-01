package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTodos(c *gin.Context) {
	cookie, err := c.Cookie("sessionID")
	if err != nil {
		handleError(c, "failed to get session", http.StatusBadRequest)
		return
	}

	s, err := h.SessionManager.GetSession(cookie)
	if err != nil {
		handleError(c, "failed to get session", http.StatusBadRequest)
		return
	}

	todos, err := h.Service.GetTodos(c, s.ID)
	if err != nil {
		handleError(c, "failed to get todos", http.StatusBadRequest)
		return
	}

	var titles []string
	for _, t := range todos {
		titles = append(titles, t.Title)
	}

	c.HTML(http.StatusOK, "profile_todo_info", gin.H{
		"Titles": titles,
	})
}
