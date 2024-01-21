package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	supa "github.com/nedpals/supabase-go"
)

type Handler struct {
	db *supa.Client
}

func NewHandler(db *supa.Client) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) InitRoute() *gin.Engine {

	router := gin.Default()
	router.Static("./frontend/static", "./frontend/static")
	router.Static("./frontend/img", "./frontend/img")

	router.LoadHTMLGlob("./frontend/template/**/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", gin.H{})
	})

	router.GET("/user-info", h.FetchUser)

	if err := router.Run(":8000"); err != nil {
		panic(err)
	}

	return router
}
