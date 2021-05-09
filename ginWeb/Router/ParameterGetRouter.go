package Router

import (
	"demo_for_gin/Controller"

	"github.com/gin-gonic/gin"
)

func loadParamterGet(e *gin.Engine) {
	paramterGetGroup := e.Group("/parameterGet")
	{
		paramterGetGroup.GET("/GetQueryString", Controller.QueryStringParameterHandler)
		paramterGetGroup.POST("/PostForm", Controller.FormParameterHandler)
		paramterGetGroup.POST("/Json", Controller.JsonParameterHandler)
	}
}
