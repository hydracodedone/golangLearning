package Router

import (
	"demo_for_gin/Controller"

	"github.com/gin-gonic/gin"
)

func loadSession(e *gin.Engine) {
	sessionGroup := e.Group("/Session")
	{
		sessionGroup.GET("/SetSession", Controller.SetSessionHandler)
		sessionGroup.GET("/GetSession", Controller.GetSessionHandler)
		sessionGroup.GET("/SetSession2", Controller.SetSessionHandler2)
		sessionGroup.GET("/GetSession2", Controller.GetSessionHandler2)
	}
}
