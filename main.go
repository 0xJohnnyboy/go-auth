package main

import (
	"github.com/joho/godotenv"
	"goauth/cmd"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}
