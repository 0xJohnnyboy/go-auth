package api

import (
	"github.com/gin-gonic/gin"
	"goauth/internal/storage"
)

var port = ":9888"
var certFile = "cert.pem"
var keyFile = "key.pem"

func Serve() error {
	r := gin.Default()
	db, err := storage.Connect()

	if err != nil {
		panic(err)
	}

	router := NewRouter(db)
	router.RegisterRoutes(r)

	return r.RunTLS(port, certFile, keyFile)
}
