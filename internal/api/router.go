package api

import (
	"github.com/gin-gonic/gin"
	. "goauth/internal/auth"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/hc", HealthCheckHandler)
	{
		unprotected := router.Group("/")
		unprotected.POST("/register", RegisterHandler)
		unprotected.POST("/login", LoginHandler)
	}

	{
		api := router.Group("/api")
		api.Use(RequireAuth())

		api.GET("/test", TestHandler)
		api.POST("/logout", LogoutHandler)
	}
}
