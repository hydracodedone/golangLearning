package Router

import (
	"demo_for_gin/Controller"
	"demo_for_gin/Middleware"
	"demo_for_gin/Utils"

	"github.com/gin-gonic/gin"
)

func loadAsync(e *gin.Engine) {
	asyncGroup := e.Group("/Async")
	{
		asyncGroup.GET("/AsyncRun", Utils.RegisterMiddleware(Controller.AsyncHandler, Middleware.LocalMiddlewareForAsync(), Middleware.LocalMiddlewareForAsync2())...)
	}
}
