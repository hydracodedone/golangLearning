package Router

import (
	"demo_for_gin/Controller"

	"github.com/gin-gonic/gin"
)

func loadLog(e *gin.Engine) {
	peopleGroup := e.Group("/Log")
	{
		peopleGroup.GET("/", Controller.LogHandler)
	}
}
