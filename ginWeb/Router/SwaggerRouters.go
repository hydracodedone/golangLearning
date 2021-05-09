package Router

import (
	"demo_for_gin/Controller"
	_ "demo_for_gin/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func loadSwagger(e *gin.Engine) {
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	SwaggerDemoGroup := e.Group("/SwaggerDemo")
	{
		SwaggerDemoGroup.GET("/Demo1", Controller.SwaggerDemoHandler)
	}
}
