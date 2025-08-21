package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goauth/internal/storage"
)

var certFile = "cert.pem"
var keyFile = "key.pem"

func Serve(port int) error {
	r := gin.Default()
	db, err := storage.Connect()

	if err != nil {
		panic(err)
	}

	router := NewRouter(db)
	router.RegisterRoutes(r)

	return r.RunTLS(fmt.Sprintf(":%d", port), certFile, keyFile)
}
