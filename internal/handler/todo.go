package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/v7ktory/htmx+go/internal/model"
)

func (h *Handler) GetTodos(c *gin.Context) {
	cookie, err := c.Cookie("sessionID")
	if err != nil {
		handleError(c, "failed to get session", http.StatusBadRequest)
		return
	}

	s, err := h.SessionManager.GetSession(c, cookie)
	if err != nil {
		handleError(c, "failed to get session", http.StatusBadRequest)
		return
	}

	todos, err := h.Service.GetTodos(c, s.ID)
	if err != nil {
		handleError(c, "failed to get todos", http.StatusBadRequest)
		return
	}

	var todoInfo []model.TodoInfo
	for _, t := range todos {
		TodoInfo := model.TodoInfo{
			ID:    t.ID,
			Title: t.Title,
		}
		todoInfo = append(todoInfo, TodoInfo)
	}

	c.HTML(http.StatusOK, "profile_todo_info", gin.H{
		"Todo_Info": todoInfo,
	})
}
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

	e, err := h.SessionManager.GetSession(c, cookie)
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

func (h *Handler) DeleteTodo(c *gin.Context) {
	cookie, err := c.Cookie("sessionID")
	if err != nil {
		handleError(c, "failed to get session", http.StatusBadRequest)
		return
	}

	s, err := h.SessionManager.GetSession(c, cookie)
	if err != nil {
		handleError(c, "failed to get session", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, "failed to parse id", http.StatusBadRequest)
		return
	}
	h.Service.DeleteTodo(c, s.ID, int32(id))

	c.Status(http.StatusOK)
}
