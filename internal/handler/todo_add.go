package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/v7ktory/htmx+go/internal/model"
)

func (h *Handler) AddTodo(c *gin.Context) {
	var todo model.TodoReq

	if err := c.ShouldBind(&todo); err != nil {
		handleError(c, "failed to parse request body", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(todo); err != nil {
		handleError(c, "failed to validate todo", http.StatusBadRequest)
		return
	}

	cookie, err := c.Cookie("sessionID")
	if err != nil {
		handleError(c, "failed to get session", http.StatusBadRequest)
		return
	}

	e, err := h.SessionManager.GetSession(cookie)
	if err != nil {
		handleError(c, "failed to get session", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateTodo(c, todo.Title, todo.Description, e.Email); err != nil {
		handleError(c, "failed to create todo", http.StatusBadRequest)
		return
	}

	c.Header("HX-Refresh", "true")
	c.Status(http.StatusCreated)
}
