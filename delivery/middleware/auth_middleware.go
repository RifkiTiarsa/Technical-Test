package middleware

import (
	"strings"
	"test-mnc/usecase"

	"github.com/gin-gonic/gin"
)

type authMiddleware struct {
	blacklistUc usecase.BlacklistUsecase
}

type AuthMiddleware interface {
	Middleware(c *gin.Context)
}

func (a *authMiddleware) Middleware(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(accessToken, "Bearer ")
	if tokenString == accessToken {
		c.JSON(401, gin.H{"error": "Bearer token is required"})
		c.Abort()
		return
	}

	claims, err := a.blacklistUc.ValidateAndProcessToken(tokenString)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("name", claims.Name)
	c.Set("email", claims.Email)
	c.Next()
}

func NewAuthMiddleware(blacklistUc usecase.BlacklistUsecase) *authMiddleware {
	return &authMiddleware{blacklistUc: blacklistUc}
}
