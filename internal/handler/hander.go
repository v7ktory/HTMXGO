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
		auth.POST("/signout", h.SignOut)
	}

	authenticated := router.Group("/profile", h.AuthMiddleware())
	{
		authenticated.GET("/my", h.UserProfile)
		authenticated.POST("/add-todo", h.AddTodo)
		authenticated.GET("/todo-info", h.GetTodos)
		authenticated.PATCH("/update-todo/:id", h.UpdateTodo)
		authenticated.DELETE("/delete-todo/:id", h.DeleteTodo)
	}

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
		userSession, err := h.SessionManager.GetSession(c, sessionID)
		if err != nil {
			handleError(c, "failed to get user session", http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Set("userSession", userSession)
		c.Next()
	}
}
