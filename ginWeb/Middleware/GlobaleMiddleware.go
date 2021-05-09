package Middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func GlobalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		beginTime := time.Now().UTC()
		c.Set("beiginTime", beginTime)
		c.Next()
		endTime := time.Now().UTC()
		executedTime := time.Since(beginTime)
		fmt.Printf("---request begin [%v] end [%v],used [%v]---\n", beginTime, endTime, executedTime)
	}
}
