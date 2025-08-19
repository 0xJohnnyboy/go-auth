package api

import (
	"github.com/gin-gonic/gin"
)

var port = ":9898"

func Serve() error {
	r := gin.Default()
	registerRoutes(r)
	return r.Run(port)
}

func registerRoutes(r *gin.Engine) {
	// r.GET("/", TestHandler)
	r.POST("/register", RegisterHandler)
	r.POST("/login", LoginHandler)
	r.GET("/logout", LogoutHandler)
}
