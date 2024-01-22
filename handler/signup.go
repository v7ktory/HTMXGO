package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nedpals/supabase-go"
	"github.com/v7ktory/htmx+go/model"
	"github.com/v7ktory/htmx+go/pkg/hash"
)

func (h *Handler) Signup(c *gin.Context) {

	var user model.User

	passwordHash, err := hash.HashPassword(user.Password)
	if err != nil {
		handleError(c, "failed to hash password", http.StatusInternalServerError)
		return
	}
	_, err = h.db.Auth.SignUp(c.Request.Context(), supabase.UserCredentials{
		Email:    user.Email,
		Password: passwordHash,
		Data:     user.Name,
	})
	if err != nil {
		handleError(c, "failed to sign up", http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}
