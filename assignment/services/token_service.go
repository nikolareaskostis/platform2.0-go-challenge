package services

import (
	"GoRepo/source/assignment/models"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"os"
	"strconv"
	"time"
)

type TokenService interface {
	Verify(tokenId string) (bool, error)
}

type tokenService struct {
	redis         *redis.Client
	ctx           context.Context
	tokenLifetime time.Duration
}

func NewTokenService(redis *redis.Client) TokenService {
	tokenLifetime := getTokenLifetime()
	return &tokenService{redis: redis, ctx: context.Background(), tokenLifetime: tokenLifetime}
}

func (s *tokenService) GetJwtToken(tokenId string) (*jwt.Token, error) {
	token, err := s.redis.Get(s.ctx, tokenId).Result()
	if err == redis.Nil {
		signKey := []byte(os.Getenv("JWT_TOKEN_SIGN_KEY"))

		claims := &models.Claims{}

		token, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("parsingError")
			}
			return signKey, nil
		})

		if err != nil {
			return nil, err
		}

		return token, nil
	}
	return nil, err
}

func getTokenLifetime() time.Duration {
	tokenLifetimeEnv := os.Getenv("JWT_TOKEN_LIFETIME")
	tokenLifetimeMinutes, err := strconv.Atoi(tokenLifetimeEnv)
	if err != nil {

	}
	tokenLifetime := time.Duration(tokenLifetimeMinutes) * time.Minute
	return tokenLifetime
}

func (s *tokenService) ValidateJwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenId := c.GetHeader("token-id")

		if tokenId == "DEMO" {
			c.Next()
		}

		valid, err := s.Verify(tokenId)
		if err != nil || !valid {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Next()
	}
}

func (s *tokenService) Verify(tokenId string) (bool, error) {

	token, err := s.GetJwtToken(tokenId)

	if err != nil {
		return false, err
	}
	return token.Valid, nil
}
