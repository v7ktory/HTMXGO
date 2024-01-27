package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/v7ktory/htmx+go/internal/service"
	"github.com/v7ktory/htmx+go/pkg/session"
)

type Handler struct {
	SessionManager *session.SessionManager
	Service        *service.Service
}

func NewHandler(sessionManager *session.SessionManager, service *service.Service) *Handler {
	return &Handler{
		SessionManager: sessionManager,
		Service:        service,
	}
}
func (h *Handler) InitRoute() *gin.Engine {

	router := gin.Default()
	router.Static("./ui/static", "./ui/static")
	router.Static("./ui/img", "./ui/img")
	router.LoadHTMLGlob("./ui/template/**/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", gin.H{})
	})
	router.GET("/user-info", h.UserInfo)

	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.SignUp)
		auth.POST("/login", h.Login)
	}

	authenticated := router.Use(h.AuthMiddleware())
	authenticated.GET("/profile/", h.UserProfile)
	return router
}

func (h *Handler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("sessionID")
		if err != nil {
			handleError(c, "failed to get sessionID cookie", http.StatusUnauthorized)
			c.Abort()
			return
		}
		userSession, err := h.SessionManager.GetSession(sessionID)
		if err != nil {
			handleError(c, "failed to get user session", http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Set("userSession", userSession)
		c.Next()
	}
}
