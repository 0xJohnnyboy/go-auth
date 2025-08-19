package api

import (
	"github.com/gin-gonic/gin"
)

var port = ":9888"

func Serve() error {
	r := gin.Default()
	RegisterRoutes(r)
	return r.Run(port)
}
