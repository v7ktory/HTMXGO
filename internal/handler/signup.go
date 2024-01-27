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
		handleError(c, "failed to parse request body", 400)
		return
	}

	// Validate the user struct
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		handleError(c, "failed to validate user", 400)
		return
	}

	if !validation.IsEmailValid(user.Email) {
		handleError(c, "invalid email", 400)
		return
	}

	// Create the user
	u, err := h.Service.Signup(user.Name, user.Email, user.Password)
	if err != nil {
		handleError(c, "failed to create user", 500)
		return
	}

	// Generate the user's session
	sessionID, err := h.SessionManager.GenerateSession(model.UserSession{
		ID:    u.ID,
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		handleError(c, "failed to generate session", 500)
		return
	}

	c.SetCookie("sessionID", sessionID, 3600, "/", "", false, true)
	c.Header("HX-Redirect", "/profile")

	c.Status(http.StatusCreated)
}