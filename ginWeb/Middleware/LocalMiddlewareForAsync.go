package Middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LocalMiddlewareForAsync() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("only for AsyncRun")
		c.Next()
	}
}

func LocalMiddlewareForAsync2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("only for AsyncRun2")
		c.Next()
	}
}
