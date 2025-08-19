package main

import (
	"github.com/joho/godotenv"
	"goauth/internal/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	api.Serve()
}
