package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// todo think
func (h *Handler) DeleteTodo(c *gin.Context) {
	cookie, err := c.Cookie("sessionID")
	if err != nil {
		handleError(c, "failed to get session", http.StatusBadRequest)
		return
	}

	session, err := h.SessionManager.GetSession(cookie)
	if err != nil {
		handleError(c, "failed to get session", http.StatusBadRequest)
		return
	}

	todos, err := h.Service.GetTodos(c, session.ID)
	if err != nil {
		handleError(c, "failed to get todos", http.StatusBadRequest)
		return
	}

	count := len(todos)

	if err := h.Service.DeleteTodo(c, int32(count)); err != nil {
		handleError(c, "failed to delete todo", http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}
