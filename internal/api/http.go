package api

import (
	"fmt"
	"github.com/gin-gonic/gin"

	. "goauth/internal/auth"

	"gorm.io/gorm"
)

type Handlers struct {
	authService *AuthService
}

func NewHandlers(db *gorm.DB) *Handlers {
	return &Handlers{
		authService: NewAuthService(db),
	}
}

func (h *Handlers) HealthCheckHandler(c *gin.Context) {
	c.String(200, "Running")
}

func (h *Handlers) TestHandler(c *gin.Context) {
	c.String(200, "Test")
}

type UserRegisterInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handlers) RegisterHandler(c *gin.Context) {
	var input UserRegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := h.authService.Register(input.Username, input.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(500, gin.H{"error": "User created but token generation failed"})
		return
	}

	refreshToken, err := h.authService.CreateRefreshToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "User created but refresh token generation failed"})
		return
	}

	c.SetCookie("token", token, 3600*24, "/", "", true, true)
	c.SetCookie("refresh_token", refreshToken, 3600*24*7, "/", "", true, true)

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

func (h *Handlers) LoginHandler(c *gin.Context) {
	var input UserLoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := h.authService.Login(input.Username, input.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := h.authService.CreateRefreshToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "User created but refresh token generation failed"})
		return
	}
	token, err := GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(500, gin.H{"error": "Token generation failed"})
		return
	}

	c.SetCookie("token", token, 3600*24, "/", "", true, true)
	c.SetCookie("refresh_token", refreshToken, 3600*24*7, "/", "", true, true)

	c.JSON(200, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
		},
	})
}

func (h *Handlers) LogoutHandler(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err == nil && refreshToken != "" {
		h.authService.RevokeRefreshToken(refreshToken)
	}

	c.SetCookie("token", "", -1, "/", "", true, true)
	c.SetCookie("refresh_token", "", -1, "/", "", true, true)

	c.JSON(200, gin.H{"message": "Logged out"})
}

func (h *Handlers) RefreshTokenHandler(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")

	fmt.Println(refreshToken)
	if err != nil {
		c.JSON(401, gin.H{"error": "No refresh token"})
		return
	}

	user, err := h.authService.ValidateRefreshToken(refreshToken)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid refresh token"})
		return
	}

	newJWT, err := GenerateToken(user.ID, user.Username)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("token", newJWT, 3600*24, "/", "", true, true)
	c.JSON(200, gin.H{"message": "Token refreshed"})
}
