package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/v7ktory/htmx+go/internal/model"
	"github.com/v7ktory/htmx+go/pkg/validation"
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

	if !validation.IsEmailValid(user.Email) {
		handleError(c, "Неверная почта", http.StatusBadRequest)
		return
	}

	usr := model.User{
		Email:    user.Email,
		Password: user.Password,
	}
	// Login the user
	sessionID, err := h.Service.Login(c, &usr)
	if err != nil {
		handleError(c, "Неверная почта или пароль", http.StatusBadRequest)
		return
	}

	c.SetCookie("sessionID", sessionID, 3600, "/", "", false, true)
	c.Header("HX-Redirect", "/profile/my")
	c.Status(http.StatusOK)
}
