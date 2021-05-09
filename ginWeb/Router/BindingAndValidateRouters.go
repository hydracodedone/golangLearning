package Router

import (
	"demo_for_gin/Controller"

	"github.com/gin-gonic/gin"
)

func loadBinding(e *gin.Engine) {
	bindingGroup := e.Group("/Binding")
	{
		bindingGroup.POST("/BindingJson", Controller.BindingJsonHandler)
		bindingGroup.POST("/BindingForm", Controller.BindingFormHandler)
	}
	validateGroup := e.Group("/Validate")
	{
		validateGroup.POST("/ValidateJson", Controller.ValidateJsonHandler)
	}
}
