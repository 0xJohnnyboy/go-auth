package api

import (
	"github.com/gin-gonic/gin"

	. "goauth/internal/auth"
)

func HealthCheckHandler(c *gin.Context) {
	c.String(200, "Running")
}

func TestHandler(c *gin.Context) {
	c.String(200, "Test")
}

type UserRegisterInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func RegisterHandler(c *gin.Context) {
	var input UserRegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := Register(input.Username, input.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

type UserLoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func LoginHandler(c *gin.Context) {
	var input UserLoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := Login(input.Username, input.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

func LogoutHandler(c *gin.Context) {
	c.String(200, "Logged out")
}
