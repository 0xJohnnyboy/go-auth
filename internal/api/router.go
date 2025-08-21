package api

import (
	"github.com/gin-gonic/gin"
	. "goauth/internal/auth"
	"gorm.io/gorm"
)
type Router struct {
	h *Handlers
}


func NewRouter(db *gorm.DB) *Router {
	return &Router{
		h: NewHandlers(db),
	}
}

func (r *Router) RegisterRoutes(router *gin.Engine) {
	router.GET("/hc", r.h.HealthCheckHandler)
	{
		unprotected := router.Group("/")
		unprotected.POST("/register", r.h.RegisterHandler)
		unprotected.POST("/login", r.h.LoginHandler)
	}

	{
		api := router.Group("/api")
		api.Use(RequireAuth())

		api.GET("/test", r.h.TestHandler)
		api.POST("/logout", r.h.LogoutHandler)
	}
}
