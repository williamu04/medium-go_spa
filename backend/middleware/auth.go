package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/williamu04/medium-clone/pkg"
)

type AuthMiddleware struct {
	logger *pkg.Logger
	jwt    *pkg.JWTGen
}

func NewAuthMiddleware(logger *pkg.Logger, jwt *pkg.JWTGen) *AuthMiddleware {
	return &AuthMiddleware{logger: logger, jwt: jwt}
}

func (m *AuthMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if authHeader := c.GetHeader("Authorization"); authHeader == "" {
			m.logger.Error("Auth header missing")
			pkg.Error(c, http.StatusUnauthorized, "Auth header required")
			c.Abort()
			return
		} else {
			if parts := strings.Split(authHeader, " "); len(parts) != 2 || parts[0] != "Bearer" {
				m.logger.Error("Invalid authorization header format")
				pkg.Error(c, http.StatusUnauthorized, "Invalid auth header format")
				c.Abort()
				return
			} else {
				token := parts[1]
				claims, err := m.jwt.Validate(token)
				if err != nil {
					m.logger.Errorf("Invalid or expired token - %v", err)
					pkg.Error(c, http.StatusUnauthorized, "Invalid or expired token")
					c.Abort()
					return
				}

				userID, ok := claims["user_id"].(uint)
				if !ok {
					m.logger.Error("Invalid user_id in token")
					pkg.Error(c, http.StatusUnauthorized, "Invalid user_id in token")
					c.Abort()
					return
				}
				c.Set("user_id", userID)
			}
		}
		c.Next()
	}
}

func (m *AuthMiddleware) OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if authHeader := c.GetHeader("Authorization"); authHeader == "" {
			m.logger.Error("Auth header missing")
			c.Set("user_id", nil)
		} else {
			if parts := strings.Split(authHeader, " "); len(parts) != 2 || parts[0] != "Bearer" {
				m.logger.Error("Invalid authorization header format")
				c.Set("user_id", nil)
			} else {
				token := parts[1]
				claims, err := m.jwt.Validate(token)
				if err != nil {
					m.logger.Errorf("Invalid or expired token - %v", err)
					c.Set("user_id", nil)
				}
				userID, ok := claims["user_id"].(uint)
				if !ok {
					m.logger.Error("Invalid user_id in token")
					c.Set("user_id", nil)
				}
				c.Set("user_id", userID)
			}
		}
		c.Next()
	}
}
