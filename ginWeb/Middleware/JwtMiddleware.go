package Middleware

import (
	_ "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
