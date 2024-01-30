package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/v7ktory/htmx+go/internal/model"
	"github.com/v7ktory/htmx+go/pkg/validation"
)

func (h *Handler) SignUp(c *gin.Context) {
	// Get the info from the request body
	var user model.SignupReq

	// Parse request body
	if err := c.ShouldBind(&user); err != nil {
		handleError(c, "failed to parse request body", http.StatusBadRequest)
		return
	}

	// Validate the user struct
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		handleError(c, "failed to validate user", http.StatusBadRequest)
		return
	}

	if !validation.IsEmailValid(user.Email) {
		handleError(c, "Неверная почта", http.StatusBadRequest)
		return
	}

	// Create the user
	u, err := h.Service.Signup(user.Name, user.Email, user.Password)
	if err != nil {
		handleError(c, "Почта уже используется", http.StatusBadRequest)
		return
	}

	// Generate the user's session
	sessionID, err := h.SessionManager.GenerateSession(model.UserSession{
		ID:    u.ID,
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		handleError(c, "failed to generate session", http.StatusInternalServerError)
		return
	}

	c.SetCookie("sessionID", sessionID, 3600, "/", "", false, true)
	c.Header("HX-Redirect", "/profile")
	c.Status(http.StatusCreated)
}
