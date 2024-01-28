package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/v7ktory/htmx+go/internal/model"
)

func (h *Handler) Login(c *gin.Context) {
	var user model.LoginReq

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

	// Login the user
	sessionID, err := h.Service.Login(user.Email, user.Password)
	if err != nil {
		handleError(c, "failed to login", http.StatusInternalServerError)
		return
	}

	c.SetCookie("sessionID", sessionID, 3600, "/", "", false, true)
	c.Header("HX-Redirect", "/profile")
	c.Status(http.StatusOK)
}
