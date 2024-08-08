package Router

import (
	"demo_for_gin/Controller"

	"github.com/gin-gonic/gin"
)

func loadCookie(e *gin.Engine) {
	peopleGroup := e.Group("/Cookie")
	{
		peopleGroup.GET("/SetCookie", Controller.SetCookieHandler)
		peopleGroup.GET("/GetCookie", Controller.GetCookieHandler)
	}
}
