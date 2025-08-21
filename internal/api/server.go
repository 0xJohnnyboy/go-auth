package api

import (
	"github.com/gin-gonic/gin"
)

var port = ":9888"
var certFile = "cert.pem"
var keyFile = "key.pem"

func Serve() error {
	r := gin.Default()
	RegisterRoutes(r)
	return r.RunTLS(port, certFile, keyFile)
}
