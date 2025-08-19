package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
	"strings"
)

var secret = os.Getenv("JWT_SECRET")

func GenerateToken(userID uint, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrTokenInvalidClaims
}

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", uint(claims["user_id"].(float64)))
		c.Set("username", claims["username"].(string))

		c.Next()
	}
}
