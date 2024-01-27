package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/v7ktory/htmx+go/internal/model"
)

func (h *Handler) Login(c *gin.Context) {

	var user model.SignInReq

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

	// Login the user
	sessionID, err := h.Service.Login(user.Email, user.Password)
	if err != nil {
		handleError(c, "failed to sign in", 500)
		return
	}

	// Set the session id as a header
	c.Writer.Header().Set("Authorization", fmt.Sprintf("Bearer %s", sessionID))

	c.JSON(200, gin.H{"success": true})
}
