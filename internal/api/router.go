package api

import (
	"github.com/gin-gonic/gin"
	. "goauth/internal/auth"
	"goauth/internal/storage"
)

func RegisterRoutes(router *gin.Engine) {
	db, err := storage.Connect()

	if err != nil {
		panic(err)
	}

	h := NewHandlers(db)

	router.GET("/hc", h.HealthCheckHandler)
	{
		unprotected := router.Group("/")
		unprotected.POST("/register", h.RegisterHandler)
		unprotected.POST("/login", h.LoginHandler)
	}

	{
		api := router.Group("/api")
		api.Use(RequireAuth())

		api.GET("/test", h.TestHandler)
		api.POST("/logout", h.LogoutHandler)
	}
}
