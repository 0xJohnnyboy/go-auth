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

	token, err := GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(500, gin.H{"error": "User created but token generation failed"})
		return
	}

	c.SetCookie(
		"token",
		token,
		3600*24,
		"/",
		"",
		true,
		true,
	)

	c.JSON(200, gin.H{
		"message": "Register successful",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
		},
	})
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

	token, err := GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(500, gin.H{"error": "Token generation failed"})
		return
	}

	c.SetCookie(
		"token",
		token,
		3600*24, // one day
		"/",
		"",
		true,
		true,
	)

	c.JSON(200, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
		},
	})
}

func LogoutHandler(c *gin.Context) {
	c.GetString("username")
	c.SetCookie("token", "", -1, "/", "", true, true)
	c.JSON(200, gin.H{"message": "Logged out"})
}
