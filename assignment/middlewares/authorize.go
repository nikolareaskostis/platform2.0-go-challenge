package middlewares

import (
	"GoRepo/source/assignment/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthMiddleware interface {
	ValidateJwtToken() gin.HandlerFunc
}

type authMiddleware struct {
	tokenService services.TokenService
}

func NewMiddleware(tokenService services.TokenService) AuthMiddleware {
	return &authMiddleware{tokenService: tokenService}
}

func (a *authMiddleware) ValidateJwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenId := c.GetHeader("token-id")

		valid, err := a.tokenService.Verify(tokenId)
		if err != nil || !valid {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Next()
	}
}
